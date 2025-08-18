package model

type ProductOffer struct {
	ID         int64
	Price      float64
	Stock      int32
	IsActive   bool
	ProductID  int64
	SupplierID int64 // user id of supplier
}
