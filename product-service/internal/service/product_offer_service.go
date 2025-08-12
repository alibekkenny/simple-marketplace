package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/alibekkenny/simple-marketplace/product-service/internal/dto"
	"github.com/alibekkenny/simple-marketplace/product-service/internal/model"
	"github.com/alibekkenny/simple-marketplace/product-service/internal/repository"
	"github.com/go-playground/validator/v10"
)

type ProductOfferService struct {
	repo      repository.ProductOfferRepository
	validator *validator.Validate
}

func NewProductOfferService(repo repository.ProductOfferRepository, validator *validator.Validate) *ProductOfferService {
	return &ProductOfferService{repo: repo, validator: validator}
}

// rpc CreateProductOffer(CreateProductOfferRequest) returns (CreateProductOfferResponse);
func (s *ProductOfferService) CreateProductOffer(ctx context.Context, input dto.CreateProductOfferInput) (int64, error) {
	if err := s.validator.Struct(input); err != nil {
		return 0, fmt.Errorf("invalid body:\n%v", err)
	}

	productOffer := model.ProductOffer{
		Price:      input.Price,
		Stock:      input.Stock,
		IsActive:   *input.IsActive,
		ProductID:  input.ProductID,
		SupplierID: input.SupplierID,
	}

	id, err := s.repo.CreateProductOffer(ctx, &productOffer)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// rpc UpdateProductOffer(UpdateProductOfferRequest) returns (UpdateProductOfferResponse);
func (s *ProductOfferService) UpdateProductOffer(ctx context.Context, id int64, input dto.UpdateProductOfferInput) (*model.ProductOffer, error) {
	if err := s.validator.Struct(input); err != nil {
		return nil, fmt.Errorf("invalid body:\n%v", err)
	}

	productOffer, err := s.repo.FindProductOfferByID(ctx, id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, fmt.Errorf("product offer with ID %v not found", id)
		} else {
			return nil, err
		}
	}

	productOffer.Price = input.Price
	productOffer.IsActive = input.IsActive
	productOffer.Stock = input.Stock

	err = s.repo.UpdateProductOffer(ctx, productOffer)
	if err != nil {
		return nil, err
	}

	return productOffer, nil
}

// rpc DeleteProductOffer(DeleteProductOfferRequest) returns (DeleteProductOfferResponse);
func (s *ProductOfferService) DeleteProductOffer(ctx context.Context, id int64) error {
	if id == 0 {
		return model.ErrNotFound
	}

	return s.repo.DeleteProductOfferByID(ctx, id)
}

// rpc GetProductOffersByProduct(GetProductOffersByProductRequest) returns (GetProductOffersByProductResponse);
func (s *ProductOfferService) GetProductOffersByProduct(ctx context.Context, productID int64) ([]model.ProductOffer, error) {
	if productID == 0 {
		return nil, model.ErrNotFound
	}

	return s.repo.FindProductOffersByProductID(ctx, productID)
}

// rpc GetProductOffersBySupplier(GetProductOffersBySupplierRequest) returns (GetProductOffersBySupplierResponse);
func (s *ProductOfferService) GetProductOffersBySupplier(ctx context.Context, supplierID int64) ([]model.ProductOffer, error) {
	if supplierID == 0 {
		return nil, model.ErrNotFound
	}

	return s.repo.FindProductOffersBySupplierID(ctx, supplierID)
}
