package dto

// rpc CreateProductOffer(CreateProductOfferRequest) returns (CreateProductOfferResponse);
type CreateProductOfferInput struct {
	Price      float32 `validate:"required"`
	Stock      int32   `validate:"required"`
	IsActive   *bool   `validate:"required"`
	ProductID  int64   `validate:"required"`
	SupplierID int64   `validate:"required"`
}

type UpdateProductOfferInput struct {
	Price    float32 `validate:"required"`
	Stock    int32   `validate:"required"`
	IsActive bool    `validate:"required"`
}
