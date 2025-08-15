package dto

type AddToCartInput struct {
	Quantity       int32 `validate:"required"`
	ProductOfferID int64 `validate:"required"`
	UserID         int64 `validate:"required"`
}

type UpdateCartInput struct {
	UserId         int64 `validate:"required"`
	ProductOfferId int64 `validte:"required"`
	Quantity       int32 `validate:"required"`
}
