package model

import (
	"time"
)

type Order struct {
	ID              int64        `json:"id"`
	TotalAmount     float64      `json:"total_amount"`
	Status          string       `json:"status"` // new, confirmed, cancelled
	CreatedAt       time.Time    `json:"created_at"`
	UserID          int64        `json:"user_id"`
	PaymentMethod   string       `json:"payment_method"`
	ShippingAddress string       `json:"shipping_address"`
	Items           []*OrderItem `json:"items"`
}
