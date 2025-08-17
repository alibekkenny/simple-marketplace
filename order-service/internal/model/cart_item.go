package model

type CartItem struct {
	Quantity       int32 `json:"quantity"`
	ProductOfferID int64 `json:"product_offer_id"`
}
