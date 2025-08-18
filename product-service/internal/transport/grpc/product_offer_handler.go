package grpc

import (
	"context"
	"errors"

	"github.com/alibekkenny/simple-marketplace/product-service/internal/dto"
	"github.com/alibekkenny/simple-marketplace/product-service/internal/model"
	"github.com/alibekkenny/simple-marketplace/product-service/internal/service"
	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductOfferHandler struct {
	pb.UnimplementedProductOfferServiceServer
	service *service.ProductOfferService
}

func NewProductOfferHandler(service *service.ProductOfferService) *ProductOfferHandler {
	return &ProductOfferHandler{service: service}
}

// rpc CreateProductOffer(CreateProductOfferRequest) returns (CreateProductOfferResponse);
func (h *ProductOfferHandler) CreateProductOffer(ctx context.Context, req *pb.CreateProductOfferRequest) (*pb.CreateProductOfferResponse, error) {
	input := dto.CreateProductOfferInput{
		Price:      req.Price,
		Stock:      &req.Stock,
		IsActive:   &req.IsActive,
		ProductID:  req.ProductId,
		SupplierID: req.SupplierId,
	}

	id, err := h.service.CreateProductOffer(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductOfferResponse{Id: id}, nil
}

// rpc UpdateProductOffer(UpdateProductOfferRequest) returns (UpdateProductOfferResponse);
func (h *ProductOfferHandler) UpdateProductOffer(ctx context.Context, req *pb.UpdateProductOfferRequest) (*pb.UpdateProductOfferResponse, error) {
	input := dto.UpdateProductOfferInput{
		Price:    req.Price,
		Stock:    &req.Stock,
		IsActive: &req.IsActive,
	}

	productOffer, err := h.service.UpdateProductOffer(ctx, req.Id, input)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "product offer with id %v not found", req.Id)
		} else {
			return nil, err
		}
	}

	responseProductOffer := pb.ProductOffer{
		Id:         productOffer.ID,
		Price:      productOffer.Price,
		Stock:      productOffer.Stock,
		IsActive:   productOffer.IsActive,
		ProductId:  productOffer.ProductID,
		SupplierId: productOffer.SupplierID,
	}

	return &pb.UpdateProductOfferResponse{Offer: &responseProductOffer}, nil
}

// rpc DeleteProductOffer(DeleteProductOfferRequest) returns (DeleteProductOfferResponse);
func (h *ProductOfferHandler) DeleteProductOffer(ctx context.Context, req *pb.DeleteProductOfferRequest) (*pb.DeleteProductOfferResponse, error) {
	id := req.Id

	err := h.service.DeleteProductOffer(ctx, id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "product offer with id %v not found", req.Id)
		} else {
			return nil, err
		}
	}

	return &pb.DeleteProductOfferResponse{}, nil
}

// rpc GetProductOffersByProduct(GetProductOffersByProductRequest) returns (GetProductOffersByProductResponse);
func (h *ProductOfferHandler) GetProductOffersByProduct(ctx context.Context, req *pb.GetProductOffersByProductRequest) (*pb.GetProductOffersByProductResponse, error) {
	productID := req.ProductId

	offers, err := h.service.GetProductOffersByProduct(ctx, productID)
	if err != nil {
		return nil, err
	}

	var offersResponse []*pb.ProductOffer

	for _, offer := range offers {
		offerResponse := pb.ProductOffer{
			Id:         offer.ID,
			Price:      offer.Price,
			Stock:      offer.Stock,
			ProductId:  offer.ProductID,
			SupplierId: offer.SupplierID,
		}

		offersResponse = append(offersResponse, &offerResponse)
	}

	return &pb.GetProductOffersByProductResponse{Offers: offersResponse}, nil
}

// rpc GetProductOffersBySupplier(GetProductOffersBySupplierRequest) returns (GetProductOffersBySupplierResponse);
func (h *ProductOfferHandler) GetProductOffersBySupplier(ctx context.Context, req *pb.GetProductOffersBySupplierRequest) (*pb.GetProductOffersBySupplierResponse, error) {
	supplierID := req.SupplierId

	offers, err := h.service.GetProductOffersBySupplier(ctx, supplierID)
	if err != nil {
		return nil, err
	}

	var offersResponse []*pb.ProductOffer

	for _, offer := range offers {
		offerResponse := pb.ProductOffer{
			Id:         offer.ID,
			Price:      offer.Price,
			Stock:      offer.Stock,
			ProductId:  offer.ProductID,
			SupplierId: offer.SupplierID,
		}

		offersResponse = append(offersResponse, &offerResponse)
	}

	return &pb.GetProductOffersBySupplierResponse{Offers: offersResponse}, nil
}

// rpc GetProductOffer(GetProductOfferRequest) returns (GetProductOfferResponse);
func (h *ProductOfferHandler) GetProductOffer(ctx context.Context, req *pb.GetProductOfferRequest) (*pb.GetProductOfferResponse, error) {
	productOffer, err := h.service.GetProductOffer(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	responseProductOffer := pb.ProductOffer{
		Id:         productOffer.ID,
		Price:      productOffer.Price,
		Stock:      productOffer.Stock,
		IsActive:   productOffer.IsActive,
		ProductId:  productOffer.ProductID,
		SupplierId: productOffer.SupplierID,
	}

	return &pb.GetProductOfferResponse{Offer: &responseProductOffer}, nil
}
