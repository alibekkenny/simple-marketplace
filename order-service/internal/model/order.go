package model

import (
	"time"
)

type Order struct {
	ID              int64
	TotalAmount     float64
	Status          string // new, confirmed, cancelled
	CreatedAt       time.Time
	EditedAt        time.Time
	UserID          int64
	PaymentMethod   string
	ShippingAddress string
	Items           []*OrderItem
}
