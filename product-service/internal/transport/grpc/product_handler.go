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

type ProductHandler struct {
	pb.UnimplementedProductServiceServer
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	input := dto.ProductInput{
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  req.CategoryId,
	}

	id, err := h.service.CreateProduct(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{Id: id}, nil
}

func (h *ProductHandler) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	input := dto.ProductInput{
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  req.CategoryId,
	}
	id := req.Id

	err := h.service.UpdateProduct(ctx, id, input)
	if err != nil {
		return nil, err
	}

	product := pb.Product{
		Id:          id,
		Name:        input.Name,
		Description: input.Description,
	}
	return &pb.UpdateProductResponse{Product: &product}, nil
}

func (h *ProductHandler) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	id := req.Id

	err := h.service.DeleteProduct(ctx, id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "product not found")
		} else {
			return nil, err
		}
	}

	return &pb.DeleteProductResponse{}, nil
}

func (h *ProductHandler) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	id := req.Id

	product, err := h.service.GetProduct(ctx, id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "product not found")
		} else {
			return nil, err
		}
	}

	responseProduct := pb.Product{
		Id:          id,
		Name:        product.Name,
		Description: product.Description,
		CategoryId:  product.CategoryID,
	}

	return &pb.GetProductResponse{Product: &responseProduct}, nil
}

func (h *ProductHandler) ListProductsByCategory(ctx context.Context, req *pb.ListProductsByCategoryRequest) (*pb.ListProductsByCategoryResponse, error) {
	categoryID := req.CategoryId
	products, err := h.service.ListProductsByCategory(ctx, categoryID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "category not found")
		} else {
			return nil, err
		}
	}

	var responseProducts []*pb.Product
	for _, product := range products {
		responseProducts = append(responseProducts, &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			CategoryId:  categoryID,
		})
	}

	return &pb.ListProductsByCategoryResponse{Products: responseProducts}, nil
}
