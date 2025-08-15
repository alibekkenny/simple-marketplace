package grpc

import (
	"context"

	pb "github.com/alibekkenny/simple-marketplace/order-service/genproto"
	"github.com/alibekkenny/simple-marketplace/order-service/internal/dto"
	"github.com/alibekkenny/simple-marketplace/order-service/internal/service"
)

type OrderHandler struct {
	pb.UnimplementedOrderServiceServer
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// Order-related RPCs
// rpc Checkout(CheckoutRequest) returns (CheckoutResponse);
func (h *OrderHandler) Checkout(ctx context.Context, req *pb.CheckoutRequest) (*pb.CheckoutResponse, error) {
	input := dto.CheckoutInput{
		UserID:          req.UserId,
		PaymentMethod:   req.PaymentMethod,
		ShippingAddress: req.ShippingAddress,
	}

	order, err := h.service.Checkout(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.CheckoutResponse{OrderId: order.ID, Status: order.Status}, nil
}

// rpc GetOrderById(GetOrderByIdRequest) returns (GetOrderByIdResponse);
func (h *OrderHandler) GetOrderById(ctx context.Context, req *pb.GetOrderByIdRequest) (*pb.GetOrderByIdResponse, error) {
	order, err := h.service.GetOrderByID(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	return &pb.GetOrderByIdResponse{Order: mapOrderToProto(order)}, nil
}

// rpc ListOrders(ListOrdersRequest) returns (ListOrdersResponse);
func (h *OrderHandler) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := h.service.ListOrders(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.ListOrdersResponse{Orders: mapOrdersToProto(orders)}, nil
}
