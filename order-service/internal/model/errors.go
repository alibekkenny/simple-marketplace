package model

import "errors"

var (
	ErrOrderNotFound        = errors.New("order not found")
	ErrCartEmpty            = errors.New("cart is empty")
	ErrCartItemNotFound     = errors.New("cart item not found")
	ErrUserNotFound         = errors.New("user not found")
	ErrProductOfferNotFound = errors.New("product offer not found")
	ErrOrderItemNotFound    = errors.New("order item not found")
)
