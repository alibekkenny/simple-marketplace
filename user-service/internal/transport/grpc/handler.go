package grpc

import (
	"context"

	pb "github.com/alibekkenny/simple-marketplace/user-service/genproto"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/dto"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/service"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
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
		return nil, err
	}

	return &pb.RegisterResponse{UserId: id}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	input := dto.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := h.service.Login(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{Token: token}, nil
}
