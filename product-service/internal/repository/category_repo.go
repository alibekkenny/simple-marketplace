package repository

import (
	"context"

	"github.com/alibekkenny/simple-marketplace/product-service/internal/model"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *model.Category) (int, error)
	UpdateCategory(ctx context.Context, category *model.Category) error
	DeleteCategoryById(ctx context.Context, id int64) error
}
