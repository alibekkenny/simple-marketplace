package service

import (
	"context"
	"fmt"
	"time"

	"github.com/alibekkenny/simple-marketplace/user-service/internal/dto"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/model"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	repo      repository.UserRepository
	jwtKey    []byte
	validator *validator.Validate
}

func NewUserService(repo repository.UserRepository, jwtKey []byte, validator *validator.Validate) *UserService {
	return &UserService{repo: repo, jwtKey: jwtKey, validator: validator}
}

func (s *UserService) Register(ctx context.Context, input dto.RegisterInput) (int64, error) {
	err := s.validator.Struct(input)
	if err != nil {
		return 0, status.Errorf(codes.InvalidArgument, "invalid input:\n%v", err)
	}

	if !isValidRole(input.Role) {
		return 0, status.Error(codes.InvalidArgument, "invalid role")
	}

	exists, err := s.repo.ExistsByEmailOrUsername(ctx, input.Email, input.Username)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, status.Errorf(codes.AlreadyExists, "user with such email or username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	fmt.Println(hashedPassword)

	newUser := model.User{
		Email:    input.Email,
		Username: input.Password,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	return s.repo.CreateUser(ctx, &newUser)
}

func (s *UserService) Login(ctx context.Context, input dto.LoginInput) (string, error) {
	err := s.validator.Struct(input)
	if err != nil {
		return "", status.Errorf(codes.InvalidArgument, "invalid input:\n%v", err)
	}

	foundUser, err := s.repo.FindUserByEmail(ctx, input.Email)
	if err != nil {
		return "", status.Errorf(codes.NotFound, "invalid credentials")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(input.Password)); err != nil {
		return "", status.Errorf(codes.NotFound, "invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    foundUser.ID,
		"user_role":  foundUser.Role,
		"expires_at": time.Now().Add(24 * time.Hour),
		"iat":        time.Now(),
	})

	return token.SignedString(s.jwtKey)
}

func isValidRole(role string) bool {
	if role == "supplier" || role == "buyer" {
		return true
	}
	return false
}
