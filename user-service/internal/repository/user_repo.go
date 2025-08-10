package repository

import (
	"context"

	"github.com/alibekkenny/simple-marketplace/user-service/internal/model"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (int64, error)
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	ExistsByEmailOrUsername(ctx context.Context, email, username string) (bool, error)
}
