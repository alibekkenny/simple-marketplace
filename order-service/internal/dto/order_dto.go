package dto

type CheckoutInput struct {
	UserID          int64  `validate:"required"`
	PaymentMethod   string `validate:"required"`
	ShippingAddress string `validate:"required,min=3"`
}
