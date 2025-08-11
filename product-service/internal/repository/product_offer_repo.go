package repository

import (
	"context"

	"github.com/alibekkenny/simple-marketplace/product-service/internal/model"
)

type ProductOfferRepository interface {
	CreateProductOffer(ctx context.Context, productOffer *model.ProductOffer) (int64, error)
	UpdateProductOffer(ctx context.Context, productOffer *model.ProductOffer) error
	DeleteProductOfferByID(ctx context.Context, id int64) error
	FindProductOffersByProductID(ctx context.Context, productID int64) ([]model.ProductOffer, error)
	FindProductOffersBySupplierID(ctx context.Context, supplierID int64) ([]model.ProductOffer, error)
	FindProductOfferByID(ctx context.Context, id int64) (*model.ProductOffer, error)
}
