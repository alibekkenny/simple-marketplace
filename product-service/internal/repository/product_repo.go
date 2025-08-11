package repository

import (
	"context"

	"github.com/alibekkenny/simple-marketplace/product-service/internal/model"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *model.Product) (int64, error)
	UpdateProduct(ctx context.Context, product *model.Product) error
	DeleteProductByID(ctx context.Context, id int64) error
	FindProductsByCategory(ctx context.Context, categoryID int64) ([]model.Product, error)
	FindProductByID(ctx context.Context, id int64) (*model.Product, error)
}
