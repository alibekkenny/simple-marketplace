package repository

import (
	"context"

	"github.com/alibekkenny/simple-marketplace/user-service/internal/model"
)

type UserRepository interface {
	FindUserByID(ctx context.Context, id int64) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (int64, error)
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	ExistsByEmailOrUsername(ctx context.Context, email, username string) (bool, error)
}
