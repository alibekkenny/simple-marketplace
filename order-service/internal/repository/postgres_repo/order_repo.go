package postgres_repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/alibekkenny/simple-marketplace/order-service/internal/model"
)

type OrderPostgresRepository struct {
	db *sql.DB
}

func NewOrderPostgresRepository(db *sql.DB) OrderPostgresRepository {
	return OrderPostgresRepository{db: db}
}

// CreateOrder(ctx context.Context, order *model.Order) (int64, error)
func (r OrderPostgresRepository) CreateOrder(ctx context.Context, order *model.Order) error {
	const stmtOrder = `INSERT INTO orders (total_amount, status, created_at, user_id, payment_method, shipping_address)
	VALUES ($1, $2, NOW(), $3, $4, $5) RETURNING id`

	const stmtItem = `INSERT INTO order_items (order_id, price, product_offer_id, quantity)
	VALUES ($1, $2, $3, $4) RETURNING id`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err = tx.QueryRowContext(ctx, stmtOrder,
		order.TotalAmount,
		order.Status,
		order.UserID,
		order.PaymentMethod,
		order.ShippingAddress,
	).Scan(&order.ID); err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}

	for _, item := range order.Items {
		item.OrderID = order.ID

		if err = tx.QueryRowContext(ctx, stmtItem,
			item.OrderID,
			item.Price,
			item.ProductOfferID,
			item.Quantity,
		).Scan(&item.ID); err != nil {
			return fmt.Errorf("failed to insert order item: %v", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// FindOrderByID(ctx context.Context, id int64) (*model.Order, error)
func (r OrderPostgresRepository) FindOrderByID(ctx context.Context, id int64) (*model.Order, error) {
	const stmtOrder = `SELECT id, total_amount, status, created_at, user_id, payment_method, shipping_address FROM orders
	WHERE id = $1`
	const stmtItem = `SELECT id, price, quantity, product_offer_id, order_id FROM order_items WHERE order_id = $1`

	var order model.Order
	if err := r.db.QueryRowContext(ctx, stmtOrder, id).Scan(&order.ID,
		&order.TotalAmount,
		&order.Status,
		&order.CreatedAt,
		&order.UserID,
		&order.PaymentMethod,
		&order.ShippingAddress,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrOrderNotFound
		}
		return nil, fmt.Errorf("failed to find order with id %v: %w", id, err)
	}

	var items []*model.OrderItem
	rows, err := r.db.QueryContext(ctx, stmtItem, order.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderItem model.OrderItem

		if err = rows.Scan(&orderItem.ID,
			&orderItem.Price,
			&orderItem.Quantity,
			&orderItem.ProductOfferID,
			&orderItem.OrderID,
		); err != nil {
			return nil, err
		}

		items = append(items, &orderItem)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	order.Items = items

	return &order, nil
}

// FindOrdersByUserID(ctx context.Context, userID int64) ([]*model.Order, error)
func (r OrderPostgresRepository) FindOrdersByUserID(ctx context.Context, userID int64) ([]*model.Order, error) {
	const stmtOrder = `SELECT id, total_amount, status, created_at, user_id, payment_method, shipping_address FROM orders
	WHERE user_id = $1`
	const stmtItem = `SELECT id, price, quantity, product_offer_id, order_id FROM order_items WHERE order_id = $1`

	rows, err := r.db.QueryContext(ctx, stmtOrder, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find orders by user_id %d: %w", userID, err)
	}
	defer rows.Close()

	var orders []*model.Order

	for rows.Next() {
		var order model.Order
		if err := rows.Scan(
			&order.ID,
			&order.TotalAmount,
			&order.Status,
			&order.CreatedAt,
			&order.UserID,
			&order.PaymentMethod,
			&order.ShippingAddress,
		); err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}

		itemRows, err := r.db.QueryContext(ctx, stmtItem, order.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to find order items by order_id %d: %w", order.ID, err)
		}

		var items []*model.OrderItem
		for itemRows.Next() {
			var item model.OrderItem
			if err = itemRows.Scan(
				&item.ID,
				&item.Price,
				&item.Quantity,
				&item.ProductOfferID,
				&item.OrderID,
			); err != nil {
				itemRows.Close()
				return nil, fmt.Errorf("failed to find order items by order_id %d: %w", order.ID, err)
			}

			items = append(items, &item)
		}

		if err = itemRows.Err(); err != nil {
			itemRows.Close()
			return nil, fmt.Errorf("failed to find order items by order_id: %w", err)
		}
		itemRows.Close()

		order.Items = items
		orders = append(orders, &order)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to find orders by user_id: %w", err)
	}
	if len(orders) == 0 {
		return nil, model.ErrOrderNotFound
	}

	return orders, nil
}
