package postgres_repo

import (
	"context"
	"database/sql"

	"github.com/alibekkenny/simple-marketplace/product-service/internal/model"
)

type CategoryPostgresRepository struct {
	db *sql.DB
}

func NewCategoryPostgresRepository(db *sql.DB) *CategoryPostgresRepository {
	return &CategoryPostgresRepository{db: db}
}

// CreateCategory(ctx context.Context, category *model.Category) (int, error)
func (r *CategoryPostgresRepository) CreateCategory(ctx context.Context, category *model.Category) (int64, error) {
	var id int64
	stmt := `INSERT INTO product_categories(name) VALUES($1) RETURNING id`

	err := r.db.QueryRowContext(ctx, stmt, category.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CategoryPostgresRepository) FindCategories(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	stmt := `SELECT id, name FROM product_categories`

	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

// UpdateCategory(ctx context.Context, category *model.Category) error
func (r *CategoryPostgresRepository) UpdateCategory(ctx context.Context, category *model.Category) error {
	stmt := `UPDATE product_categories SET name = $1 WHERE id = $2`

	row, err := r.db.ExecContext(ctx, stmt, category.Name, category.ID)
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

// DeleteCategoryByID(ctx context.Context, id int64) error
func (r *CategoryPostgresRepository) DeleteCategoryByID(ctx context.Context, id int64) error {
	stmt := `DELETE FROM product_categories WHERE id = $1`

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
