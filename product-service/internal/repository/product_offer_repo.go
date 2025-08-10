package repository

import (
	"context"

	"github.com/alibekkenny/simple-marketplace/product-service/internal/model"
)

type ProductOffer interface {
	CreateProductOffet(ctx context.Context, productOffer *model.ProductOffer) (int, error)
	GetProductOffersByProductID(ctx context.Context, productID int64) ([]ProductOffer, error)
	GetProductOffersBySupplierID(ctx context.Context, supplierID int64) ([]ProductOffer, error)
}
