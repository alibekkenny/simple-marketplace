package dto

// rpc CreateProduct(CreateProductRequest) returns (CreateProductResponse);
type ProductInput struct {
	Name        string `validate:"required,min=3"`
	Description string `validate:"min=3"`
	CategoryID  int64  `validate:"required"`
}
