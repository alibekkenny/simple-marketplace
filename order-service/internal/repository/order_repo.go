package repository

import (
	"context"

	"github.com/alibekkenny/simple-marketplace/order-service/internal/model"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *model.Order) error
	FindOrderByID(ctx context.Context, id int64) (*model.Order, error)
	FindOrdersByUserID(ctx context.Context, userID int64) ([]*model.Order, error)
}
