package repository

import (
	"context"

	"github.com/alibekkenny/simple-marketplace/product-service/internal/model"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *model.Category) (int64, error)
	FindCategories(ctx context.Context) ([]model.Category, error)
	UpdateCategory(ctx context.Context, category *model.Category) error
	DeleteCategoryByID(ctx context.Context, id int64) error
}
