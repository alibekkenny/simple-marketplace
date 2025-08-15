package model

type OrderItem struct {
	Price          float64
	Quantity       int32
	ProductOfferID int64
	OrderID        int64
}
