package model

type OrderItem struct {
	ID             int64   `json:"id"`
	Price          float64 `json:"price"`
	Quantity       int32   `json:"quantity"`
	ProductOfferID int64   `json:"product_offer_id"`
	OrderID        int64   `json:"order_id"`
}
