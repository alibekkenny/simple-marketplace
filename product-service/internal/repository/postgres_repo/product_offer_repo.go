package postgres_repo

import (
	"context"
	"database/sql"

	"github.com/alibekkenny/simple-marketplace/product-service/internal/model"
)

type ProductOfferPostgresRepository struct {
	db *sql.DB
}

func NewProductOfferPostgresRepository(db *sql.DB) *ProductOfferPostgresRepository {
	return &ProductOfferPostgresRepository{db: db}
}

// CreateProductOffer(ctx context.Context, productOffer *model.ProductOffer) (int, error)
func (r *ProductOfferPostgresRepository) CreateProductOffer(ctx context.Context, productOffer *model.ProductOffer) (int64, error) {
	var id int64
	stmt := `INSERT INTO product_offers(price, stock, is_active, product_id, supplier_id)
	VALUES($1, $2, $3, $4, $5) RETURNING id`

	err := r.db.QueryRowContext(ctx, stmt, productOffer.Price, productOffer.Stock, productOffer.IsActive, productOffer.ProductID, productOffer.SupplierID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateProductOffer(ctx context.Context, productOffer *model.ProductOffer) error
func (r *ProductOfferPostgresRepository) UpdateProductOffer(ctx context.Context, productOffer *model.ProductOffer) error {
	stmt := `UPDATE product_offers 
	SET price = $1, stock = $2, is_active = $3, product_id = $4, supplier_id = $5 
	WHERE id = $6`

	row, err := r.db.ExecContext(ctx, stmt, productOffer.Price, productOffer.Stock, productOffer.IsActive, productOffer.ProductID, productOffer.SupplierID, productOffer.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return model.ErrNotFound
	}

	return nil
}

// DeleteProductOfferByID(ctx context.Context, id int64) error
func (r *ProductOfferPostgresRepository) DeleteProductOfferByID(ctx context.Context, id int64) error {
	stmt := `DELETE FROM product_offers WHERE id = $1`

	row, err := r.db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return model.ErrNotFound
	}

	return nil
}

// GetProductOffersByProductID(ctx context.Context, productID int64) ([]model.ProductOffer, error)
func (r *ProductOfferPostgresRepository) FindProductOffersByProductID(ctx context.Context, productID int64) ([]model.ProductOffer, error) {
	var productOffers []model.ProductOffer
	stmt := `SELECT id, price, stock, is_active, supplier_id FROM product_offers WHERE product_id = $1`

	rows, err := r.db.QueryContext(ctx, stmt, productID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var productOffer model.ProductOffer
		if err := rows.Scan(&productOffer.ID, &productOffer.Price, &productOffer.Stock, &productOffer.IsActive, &productOffer.SupplierID); err != nil {
			return nil, err
		}
		productOffer.ProductID = productID
		productOffers = append(productOffers, productOffer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return productOffers, nil
}

// GetProductOffersBySupplierID(ctx context.Context, supplierID int64) ([]model.ProductOffer, error)
func (r *ProductOfferPostgresRepository) FindProductOffersBySupplierID(ctx context.Context, supplierID int64) ([]model.ProductOffer, error) {
	var productOffers []model.ProductOffer
	stmt := `SELECT id, price, stock, is_active, product_id FROM product_offers WHERE supplier_id = $1`

	rows, err := r.db.QueryContext(ctx, stmt, supplierID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var productOffer model.ProductOffer
		if err := rows.Scan(&productOffer.ID, &productOffer.Price, &productOffer.Stock, &productOffer.IsActive, &productOffer.ProductID); err != nil {
			return nil, err
		}
		productOffer.SupplierID = supplierID
		productOffers = append(productOffers, productOffer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return productOffers, nil
}

// GetProductOfferByID(ctx context.Context, id int64) (*model.ProductOffer, error)
func (r *ProductOfferPostgresRepository) FindProductOfferByID(ctx context.Context, id int64) (*model.ProductOffer, error) {
	var productOffer model.ProductOffer
	stmt := `SELECT id, price, stock, is_active, product_id, supplier_id FROM product_offers WHERE id = $1`

	err := r.db.QueryRowContext(ctx, stmt, id).Scan(&productOffer.ID, &productOffer.Price, &productOffer.Stock,
		&productOffer.IsActive, &productOffer.ProductID, &productOffer.SupplierID)

	if err != nil {
		return nil, err
	}

	return &productOffer, nil
}
