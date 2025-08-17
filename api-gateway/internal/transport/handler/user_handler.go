package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/user"
)

type UserHandler struct {
	client pb.UserServiceClient
}

func NewUserHandler(client pb.UserServiceClient) *UserHandler {
	return &UserHandler{client: client}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	res, err := h.client.Login(r.Context(), &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed login:\n%v", err), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request:\n%v", err), http.StatusBadRequest)
		return
	}

	res, err := h.client.Register(r.Context(), &pb.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("register failed:\n%v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
