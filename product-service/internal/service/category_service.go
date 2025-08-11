package service

import (
	"context"

	"github.com/alibekkenny/simple-marketplace/product-service/internal/dto"
	"github.com/alibekkenny/simple-marketplace/product-service/internal/model"
	"github.com/alibekkenny/simple-marketplace/product-service/internal/repository"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	repo      repository.CategoryRepository
	validator *validator.Validate
}

func NewCategoryService(repo repository.CategoryRepository, validator *validator.Validate) *CategoryService {
	return &CategoryService{repo: repo, validator: validator}
}

func (s *CategoryService) CreateCategory(ctx context.Context, input dto.CategoryInput) (int64, error) {
	if err := s.validator.Struct(input); err != nil {
		return 0, status.Errorf(codes.InvalidArgument, "invalid body:\n%v", err)
	}

	category := model.Category{
		Name: input.Name,
	}

	return s.repo.CreateCategory(ctx, &category)
}

// rpc UpdateCategory(UpdateCategoryRequest) returns (UpdateCategoryResponse);
func (s *CategoryService) UpdateCategory(ctx context.Context, id int64, input dto.CategoryInput) error {
	if err := s.validator.Struct(input); err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid body:\n%v", err)
	}

	category := model.Category{
		ID:   id,
		Name: input.Name,
	}

	return s.repo.UpdateCategory(ctx, &category)
}

// rpc DeleteCategory(DeleteCategoryRequest) returns (DeleteCategoryResponse);
func (s *CategoryService) DeleteCategory(ctx context.Context, id int64) error {
	return s.repo.DeleteCategoryByID(ctx, id)
}

// rpc ListCategories(ListCategoriesRequest) returns (ListCategoriesResponse);
func (s *CategoryService) ListCategories(ctx context.Context) ([]model.Category, error) {
	return s.repo.FindCategories(ctx)
}
