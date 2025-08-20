package grpc

import (
	"context"
	"errors"

	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/user"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/app"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/dto"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/model"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	app     *app.Application
	service *service.UserService
}

func NewUserHandler(service *service.UserService, app *app.Application) *UserHandler {
	return &UserHandler{service: service, app: app}
}

func (h *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	input := dto.RegisterInput{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}

	id, err := h.service.Register(ctx, input)
	if err != nil {
		h.app.Logger.Error().
			Err(err).
			Str("method", "Register").
			Str("email", req.Email).
			Str("username", req.Username).
			Msg("failed to register user")

		if errors.Is(err, model.ErrInvalidInput) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, model.ErrDuplicate) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		} else {
			return nil, status.Error(codes.Internal, model.ErrInternal.Error())
		}
	}

	h.app.Logger.Info().
		Str("method", "Register").
		Str("email", req.Email).
		Str("username", req.Username).
		Int64("user_id", id).
		Msg("user registered successfully")

	return &pb.RegisterResponse{UserId: id}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	input := dto.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := h.service.Login(ctx, input)
	if err != nil {
		h.app.Logger.Error().
			Err(err).
			Str("method", "Login").
			Str("email", req.Email).
			Msg("failed login attempt")

		if errors.Is(err, model.ErrInvalidInput) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, model.ErrInvalidCredentials) {
			return nil, status.Error(codes.Unauthenticated, model.ErrInvalidCredentials.Error())
		} else {
			return nil, status.Error(codes.Internal, model.ErrInternal.Error())
		}
	}

	h.app.Logger.Info().
		Str("method", "Login").
		Str("email", req.Email).
		Msg("user logged in successfully")

	return &pb.LoginResponse{Token: token}, nil
}
