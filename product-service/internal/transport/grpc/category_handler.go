package grpc

import (
	"context"
	"errors"

	pb "github.com/alibekkenny/simple-marketplace/product-service/genproto"
	"github.com/alibekkenny/simple-marketplace/product-service/internal/dto"
	"github.com/alibekkenny/simple-marketplace/product-service/internal/model"
	"github.com/alibekkenny/simple-marketplace/product-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryHandler struct {
	pb.UnimplementedCategoryServiceServer
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	input := dto.CategoryInput{
		Name: req.Name,
	}

	id, err := h.service.CreateCategory(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCategoryResponse{Id: id}, nil
}

func (h *CategoryHandler) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error) {
	id := req.Id
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "field \"id\" is required")
	}
	input := dto.CategoryInput{
		Name: req.Name,
	}

	err := h.service.UpdateCategory(ctx, id, input)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "category not found")
		} else {
			return nil, err
		}
	}

	category := pb.Category{
		Id:   id,
		Name: input.Name,
	}

	return &pb.UpdateCategoryResponse{Category: &category}, nil
}

func (h *CategoryHandler) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {
	id := req.Id

	err := h.service.DeleteCategory(ctx, id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "category not found")
		} else {
			return nil, err
		}
	}

	return &pb.DeleteCategoryResponse{}, nil
}

func (h *CategoryHandler) ListCategories(ctx context.Context, req *pb.ListCategoriesRequest) (*pb.ListCategoriesResponse, error) {
	categories, err := h.service.ListCategories(ctx)
	if err != nil {
		return nil, err
	}

	var responseCategories []*pb.Category

	for _, category := range categories {
		responseCategories = append(responseCategories, &pb.Category{
			Name: category.Name,
		})
	}

	return &pb.ListCategoriesResponse{Categories: responseCategories}, nil
}
