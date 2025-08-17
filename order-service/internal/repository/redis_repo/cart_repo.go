package redis_repo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alibekkenny/simple-marketplace/order-service/internal/model"
	"github.com/redis/go-redis/v9"
)

type CartRedisRepository struct {
	client *redis.Client
}

func NewCartRedisRepository(client *redis.Client) CartRedisRepository {
	return CartRedisRepository{client: client}
}

// AddToCart(ctx context.Context, userID int64, item model.CartItem) error
func (r CartRedisRepository) AddToCart(ctx context.Context, userID int64, item *model.CartItem) error {
	key := fmt.Sprintf("cart:%d", userID)

	data, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to add to cart by user_id %d: %w", userID, err)
	}

	return r.client.HSet(ctx, key, item.ProductOfferID, data).Err()
}

// UpdateCartItem(ctx context.Context, userID int64, item model.CartItem) error
func (r CartRedisRepository) UpdateCartItem(ctx context.Context, userID int64, item *model.CartItem) error {
	key := fmt.Sprintf("cart:%d", userID)

	data, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to update cart item by user_id %d: %w", userID, err)
	}

	return r.client.HSet(ctx, key, item.ProductOfferID, data).Err()
}

// RemoveCartItem(ctx context.Context, userID, productOfferID int64) error
func (r CartRedisRepository) RemoveCartItem(ctx context.Context, userID, productOfferID int64) error {
	key := fmt.Sprintf("cart:%d", userID)
	field := fmt.Sprintf("%d", productOfferID)

	return r.client.HDel(ctx, key, field).Err()
}

// GetCart(ctx context.Context, userID int64) ([]*model.CartItem, error)
func (r CartRedisRepository) GetCart(ctx context.Context, userID int64) ([]*model.CartItem, error) {
	key := fmt.Sprintf("cart:%d", userID)

	values, err := r.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get cart by user_id %d: %w", userID, err)
	}

	items := []*model.CartItem{}
	for _, v := range values {
		var item *model.CartItem
		if err := json.Unmarshal([]byte(v), item); err != nil {
			return nil, fmt.Errorf("failed to get cart by user_id %d: %w", userID, err)
		}
		items = append(items, item)
	}

	return items, nil
}

// ClearCart(ctx context.Context, userID int64) error
func (r CartRedisRepository) ClearCart(ctx context.Context, userID int64) error {
	key := fmt.Sprintf("cart:%d", userID)

	if err := r.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to clear cart by user_id %d: %w", userID, err)
	}

	return nil
}
