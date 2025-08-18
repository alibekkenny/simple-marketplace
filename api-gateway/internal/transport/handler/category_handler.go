package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryHandler struct {
	client pb.CategoryServiceClient
}

func NewCategoryHandler(client pb.CategoryServiceClient) *CategoryHandler {
	return &CategoryHandler{client: client}
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request:\n%v", err), http.StatusBadRequest)
		return
	}

	res, err := h.client.CreateCategory(r.Context(), &pb.CreateCategoryRequest{
		Name: req.Name,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.InvalidArgument:
				http.Error(w, st.Message(), http.StatusBadRequest)
			case codes.NotFound:
				http.Error(w, st.Message(), http.StatusNotFound)
			case codes.AlreadyExists:
				http.Error(w, st.Message(), http.StatusConflict)
			default:
				http.Error(w, st.Message(), http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "unexpected error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid id:\n%v", err), http.StatusBadRequest)
		return
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request:\n%v", err), http.StatusBadRequest)
		return
	}

	res, err := h.client.UpdateCategory(r.Context(), &pb.UpdateCategoryRequest{
		Id:   id,
		Name: req.Name,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.InvalidArgument:
				http.Error(w, st.Message(), http.StatusBadRequest)
			case codes.NotFound:
				http.Error(w, st.Message(), http.StatusNotFound)
			default:
				http.Error(w, st.Message(), http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "unexpected error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid id:\n%v", err), http.StatusBadRequest)
		return
	}

	_, err = h.client.DeleteCategory(r.Context(), &pb.DeleteCategoryRequest{
		Id: id,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.InvalidArgument:
				http.Error(w, st.Message(), http.StatusBadRequest)
			case codes.NotFound:
				http.Error(w, st.Message(), http.StatusNotFound)
			default:
				http.Error(w, st.Message(), http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "unexpected error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	res, err := h.client.ListCategories(r.Context(), &pb.ListCategoriesRequest{})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.InvalidArgument:
				http.Error(w, st.Message(), http.StatusBadRequest)
			case codes.NotFound:
				http.Error(w, st.Message(), http.StatusNotFound)
			default:
				http.Error(w, st.Message(), http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "unexpected error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
