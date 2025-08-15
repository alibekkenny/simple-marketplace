package repository

import (
	"context"

	"github.com/alibekkenny/simple-marketplace/order-service/internal/model"
)

type CartRepository interface {
	AddToCart(ctx context.Context, userID int64, item model.CartItem) error
	UpdateCartItem(ctx context.Context, userID int64, item model.CartItem) error
	RemoveCartItem(ctx context.Context, userID, productOfferID int64) error
	GetCart(ctx context.Context, userID int64) ([]*model.CartItem, error)
	ClearCart(ctx context.Context, userID int64) error
}
