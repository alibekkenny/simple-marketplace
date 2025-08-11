package dto

// c CreateCategory(CreateCategoryRequest) returns (CreateCategoryResponse);
type CategoryInput struct {
	Name string `validate:"required,min=3"`
}
