package grpc

import (
	"context"

	pb "github.com/alibekkenny/simple-marketplace/order-service/genproto"
	"github.com/alibekkenny/simple-marketplace/order-service/internal/dto"
	"github.com/alibekkenny/simple-marketplace/order-service/internal/service"
)

type CartHandler struct {
	pb.UnimplementedOrderServiceServer
	service *service.CartService
}

// Cart-related RPCs
// rpc AddToCart(AddToCartRequest) returns (AddToCartResponse);
func (h *CartHandler) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	input := dto.AddToCartInput{
		ProductOfferID: req.ProductOfferId,
		Quantity:       int32(req.Quantity),
		UserID:         req.UserId,
	}

	err := h.service.AddToCart(ctx, &input)
	if err != nil {
		return nil, err
	}

	return &pb.AddToCartResponse{Success: true}, nil
}

// rpc GetCart(GetCartRequest) returns (GetCartResponse);
func (h *CartHandler) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	cartItems, err := h.service.GetCart(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.GetCartResponse{Items: mapCartItemsToProto(cartItems)}, nil
}

// rpc UpdateCartItem(UpdateCartItemRequest) returns (UpdateCartItemResponse);
func (h *CartHandler) UpdateCartItem(ctx context.Context, req *pb.UpdateCartItemRequest) (*pb.UpdateCartItemResponse, error) {
	input := dto.UpdateCartInput{
		UserId:         req.UserId,
		ProductOfferId: req.ProductOfferId,
		Quantity:       req.Quantity,
	}

	err := h.service.UpdateCartItem(ctx, &input)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateCartItemResponse{Success: true}, nil
}

// rpc RemoveCartItem(RemoveCartItemRequest) returns (RemoveCartItemResponse);
func (h *CartHandler) RemoveCartItem(ctx context.Context, req *pb.RemoveCartItemRequest) (*pb.RemoveCartItemResponse, error) {
	err := h.service.RemoveCartItem(ctx, req.UserId, req.ProductOfferId)
	if err != nil {
		return nil, err
	}

	return &pb.RemoveCartItemResponse{Success: true}, nil
}

// rpc ClearCart(ClearCartRequest) returns (ClearCartResponse);
func (h *CartHandler) ClearCart(ctx context.Context, req *pb.ClearCartRequest) (*pb.ClearCartResponse, error) {
	err := h.service.ClearCart(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.ClearCartResponse{Success: true}, nil
}
