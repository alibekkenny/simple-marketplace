package service

import (
	"context"

	"github.com/alibekkenny/simple-marketplace/product-service/internal/dto"
	"github.com/alibekkenny/simple-marketplace/product-service/internal/model"
	"github.com/alibekkenny/simple-marketplace/product-service/internal/repository"
	"github.com/go-playground/validator/v10"
)

type ProductService struct {
	validator *validator.Validate
	repo      repository.ProductRepository
}

func NewProductService(validator *validator.Validate, repo *repository.ProductRepository) *ProductService {
	return &ProductService{validator: validator, repo: *repo}
}

// rpc CreateProduct(CreateProductRequest) returns (CreateProductResponse);
func (s *ProductService) CreateProduct(ctx context.Context, input dto.ProductInput) (int64, error) {
	err := s.validator.Struct(input)
	if err != nil {
		return 0, err
	}

	product := model.Product{
		Name:        input.Name,
		Description: input.Description,
		CategoryID:  input.CategoryID,
	}

	return s.repo.CreateProduct(ctx, &product)
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int64, input dto.ProductInput) error {
	err := s.validator.Struct(input)
	if err != nil {
		return err
	}

	if id == 0 {
		return model.ErrNotFound
	}

	product := model.Product{
		ID:          id,
		Name:        input.Name,
		Description: input.Name,
		CategoryID:  input.CategoryID,
	}

	return s.repo.UpdateProduct(ctx, &product)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int64) error {
	if id == 0 {
		return model.ErrNotFound
	}

	return s.repo.DeleteProductByID(ctx, id)
}

// rpc GetProduct(GetProductRequest) returns (GetProductResponse);
func (s *ProductService) GetProduct(ctx context.Context, id int64) (*model.Product, error) {
	if id == 0 {
		return nil, model.ErrNotFound
	}

	return s.repo.FindProductByID(ctx, id)
}

// rpc ListProductsByCategory(ListProductsByCategoryRequest) returns (ListProductsByCategoryResponse);
func (s *ProductService) ListProductsByCategory(ctx context.Context, id int64) ([]model.Product, error) {
	if id == 0 {
		return nil, model.ErrNotFound
	}

	return s.repo.FindProductsByCategory(ctx, id)
}
