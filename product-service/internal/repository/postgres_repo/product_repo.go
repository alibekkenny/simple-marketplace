package postgres_repo

import (
	"context"
	"database/sql"

	"github.com/alibekkenny/simple-marketplace/product-service/internal/model"
)

type ProductPostgresRepository struct {
	db *sql.DB
}

func NewProductPostgresRepository(db *sql.DB) *ProductPostgresRepository {
	return &ProductPostgresRepository{db: db}
}

// CreateProduct(ctx context.Context, product *model.Product) (int, error)
func (r *ProductPostgresRepository) CreateProduct(ctx context.Context, product *model.Product) (int64, error) {
	var id int64
	stmt := `INSERT INTO products(name, description, category_id) 
	VALUES($1, $2, $3) RETURNING id`

	err := r.db.QueryRowContext(ctx, stmt, product.Name, product.Description, product.CategoryID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateProduct(ctx context.Context, product *model.Product) error
func (r *ProductPostgresRepository) UpdateProduct(ctx context.Context, product *model.Product) error {
	stmt := `UPDATE products SET name = $1, description = $2, category_id = $3 WHERE id = $4`

	row, err := r.db.ExecContext(ctx, stmt, product.Name, product.Description, product.CategoryID, product.ID)
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

// DeleteProductByID(ctx context.Context, id int64) error
func (r *ProductPostgresRepository) DeleteProductByID(ctx context.Context, id int64) error {
	stmt := `DELETE FROM products WHERE id = $1`

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

// GetProductsByCategory(ctx context.Context, categoryID int64) ([]model.Product, error)
func (r *ProductPostgresRepository) FindProductsByCategory(ctx context.Context, categoryID int64) ([]model.Product, error) {
	var products []model.Product
	stmt := `SELECT id, name, description FROM products WHERE category_id = $1`

	rows, err := r.db.QueryContext(ctx, stmt, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description); err != nil {
			return nil, err
		}
		product.CategoryID = categoryID

		products = append(products, product)
	}

	// check for errors during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// GetProductByID(ctx context.Context, id int64) (*model.Product, error)
func (r *ProductPostgresRepository) FindProductByID(ctx context.Context, id int64) (*model.Product, error) {
	var product model.Product
	stmt := `SELECT id, name, description, category_id FROM products WHERE id = $1`

	err := r.db.QueryRowContext(ctx, stmt, id).Scan(&product.ID, &product.Name, &product.Description, &product.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrNotFound
		}
		return nil, err
	}

	return &product, nil
}
