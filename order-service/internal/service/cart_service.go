package service

import (
	"context"

	"github.com/alibekkenny/simple-marketplace/order-service/internal/dto"
	"github.com/alibekkenny/simple-marketplace/order-service/internal/model"
	"github.com/alibekkenny/simple-marketplace/order-service/internal/repository"
	"github.com/go-playground/validator/v10"
)

type CartService struct {
	validator *validator.Validate
	repo      repository.CartRepository
}

func (s *CartService) AddToCart(ctx context.Context, input *dto.AddToCartInput) error {
	if err := s.validator.Struct(input); err != nil {
		return err
	}

	cartItem := model.CartItem{
		Quantity:       input.Quantity,
		ProductOfferID: input.ProductOfferID,
	}

	err := s.repo.AddToCart(ctx, input.UserID, cartItem)
	if err != nil {
		return err
	}

	return nil
}

func (s *CartService) GetCart(ctx context.Context, userID int64) ([]*model.CartItem, error) {
	if userID <= 0 {
		return nil, model.ErrNotFound
	}

	return s.repo.GetCart(ctx, userID)
}

func (s *CartService) UpdateCartItem(ctx context.Context, input *dto.UpdateCartInput) error {
	if err := s.validator.Struct(input); err != nil {
		return err
	}

	if input.UserId <= 0 {
		return model.ErrNotFound
	}

	cartItem := model.CartItem{
		Quantity:       input.Quantity,
		ProductOfferID: input.ProductOfferId,
	}

	return s.repo.UpdateCartItem(ctx, input.UserId, cartItem)
}

func (s *CartService) RemoveCartItem(ctx context.Context, userID, productOfferID int64) error {
	if userID <= 0 || productOfferID <= 0 {
		return model.ErrNotFound
	}

	return s.repo.RemoveCartItem(ctx, userID, productOfferID)
}

func (s *CartService) ClearCart(ctx context.Context, userID int64) error {
	if userID <= 0 {
		return model.ErrNotFound
	}

	return s.repo.ClearCart(ctx, userID)
}
