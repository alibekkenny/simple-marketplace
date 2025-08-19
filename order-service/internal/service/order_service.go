package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/alibekkenny/simple-marketplace/order-service/internal/dto"
	"github.com/alibekkenny/simple-marketplace/order-service/internal/model"
	"github.com/alibekkenny/simple-marketplace/order-service/internal/repository"
	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/product"
	"github.com/go-playground/validator/v10"
)

type OrderService struct {
	validator     *validator.Validate
	repo          repository.OrderRepository
	cartService   *CartService
	productClient pb.ProductOfferServiceClient
}

func NewOrderService(validator *validator.Validate, repo repository.OrderRepository, cartService *CartService, productOfferClient pb.ProductOfferServiceClient) *OrderService {
	return &OrderService{validator: validator, repo: repo, cartService: cartService, productClient: productOfferClient}
}

// rpc Checkout(CheckoutRequest) returns (CheckoutResponse);
func (s *OrderService) Checkout(ctx context.Context, input dto.CheckoutInput) (*model.Order, error) {
	if err := s.validator.Struct(input); err != nil {
		return nil, err
	}

	items, err := s.getCartItems(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	orderItems, totalAmount, err := s.prepareOrderItems(ctx, items)
	if err != nil {
		return nil, err
	}

	order := s.buildOrder(totalAmount, input, orderItems)

	if err := s.saveOrderWithOrderItems(ctx, &order); err != nil {
		return nil, err
	}

	if err := s.cartService.ClearCart(ctx, input.UserID); err != nil {
		return nil, err
	}

	return &order, nil
}

// rpc GetOrderById(GetOrderByIdRequest) returns (GetOrderByIdResponse);
func (s *OrderService) GetOrderByID(ctx context.Context, id int64) (*model.Order, error) {
	if id <= 0 {
		return nil, model.ErrOrderNotFound
	}

	return s.repo.FindOrderByID(ctx, id)
}

// rpc ListOrders(ListOrdersRequest) returns (ListOrdersResponse);
func (s *OrderService) ListOrders(ctx context.Context, userID int64) ([]*model.Order, error) {
	if userID <= 0 {
		return nil, model.ErrOrderNotFound
	}

	return s.repo.FindOrdersByUserID(ctx, userID)
}

func (s *OrderService) getCartItems(ctx context.Context, userID int64) ([]*model.CartItem, error) {
	items, err := s.cartService.GetCart(ctx, userID)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, errors.New("cart is empty")
	}

	return items, nil
}

func (s *OrderService) prepareOrderItems(ctx context.Context, items []*model.CartItem) ([]*model.OrderItem, float64, error) {
	var totalAmount float64
	orderItems := make([]*model.OrderItem, 0, len(items))

	for _, item := range items {
		productOffer, err := s.productClient.GetProductOffer(ctx, &pb.GetProductOfferRequest{Id: item.ProductOfferID})
		if err != nil {
			return nil, 0, err
		}

		price := productOffer.Offer.Price
		totalAmount += (price * float64(item.Quantity))

		orderItems = append(orderItems, &model.OrderItem{
			Price:          price,
			Quantity:       item.Quantity,
			ProductOfferID: item.ProductOfferID,
		})
	}

	return orderItems, totalAmount, nil
}

func (s *OrderService) buildOrder(totalAmount float64, input dto.CheckoutInput, orderItems []*model.OrderItem) model.Order {
	return model.Order{
		TotalAmount:     totalAmount,
		Status:          "new",
		UserID:          input.UserID,
		PaymentMethod:   input.PaymentMethod,
		ShippingAddress: input.ShippingAddress,
		Items:           orderItems,
	}
}

func (s *OrderService) saveOrderWithOrderItems(ctx context.Context, order *model.Order) error {
	if err := s.repo.CreateOrder(ctx, order); err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}

	return nil
}
