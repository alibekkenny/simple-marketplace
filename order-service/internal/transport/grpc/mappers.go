package grpc

import (
	pb "github.com/alibekkenny/simple-marketplace/order-service/genproto"
	"github.com/alibekkenny/simple-marketplace/order-service/internal/model"
)

func mapOrdersToProto(orders []*model.Order) []*pb.Order {
	responseOrders := make([]*pb.Order, len(orders))
	for i, order := range orders {
		responseOrders[i] = mapOrderToProto(order)
	}

	return responseOrders
}

func mapOrderToProto(order *model.Order) *pb.Order {
	return &pb.Order{
		Id:        order.ID,
		UserId:    order.UserID,
		Status:    order.Status,
		CreatedAt: order.CreatedAt.String(),
		Items:     mapOrderItemsToProto(order.Items),
	}
}

func mapOrderItemsToProto(orderItems []*model.OrderItem) []*pb.OrderItem {
	responseOrderItems := make([]*pb.OrderItem, len(orderItems))
	for i, orderItem := range orderItems {
		responseOrderItems[i] = mapOrderItemToProto(orderItem)
	}

	return responseOrderItems
}

func mapOrderItemToProto(orderItem *model.OrderItem) *pb.OrderItem {
	return &pb.OrderItem{
		ProductOfferId: orderItem.ProductOfferID,
		Quantity:       orderItem.Quantity,
		Price:          float64(orderItem.Price),
	}
}

func mapCartItemsToProto(cartItems []*model.CartItem) []*pb.CartItem {
	items := make([]*pb.CartItem, len(cartItems))

	for i, cartItem := range cartItems {
		items[i] = mapCartItemToProto(cartItem)
	}

	return items
}

func mapCartItemToProto(cartItem *model.CartItem) *pb.CartItem {
	return &pb.CartItem{
		ProductOfferId: cartItem.ProductOfferID,
		Quantity:       cartItem.Quantity,
	}
}
