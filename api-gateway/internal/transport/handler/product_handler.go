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

type ProductHandler struct {
	client pb.ProductServiceClient
}

func NewProductHandler(client pb.ProductServiceClient) *ProductHandler {
	return &ProductHandler{client: client}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		CategoryID  int64  `json:"category_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request:\n%v", err), http.StatusBadRequest)
		return
	}

	res, err := h.client.CreateProduct(r.Context(), &pb.CreateProductRequest{
		Name:        req.Name,
		Description: req.Description,
		CategoryId:  req.CategoryID,
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

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid id:\n%v", err), http.StatusBadRequest)
		return
	}

	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		CategoryID  int64  `json:"category_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request:\n%v", err), http.StatusBadRequest)
		return
	}

	res, err := h.client.UpdateProduct(r.Context(), &pb.UpdateProductRequest{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
		CategoryId:  req.CategoryID,
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

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid id:\n%v", err), http.StatusBadRequest)
		return
	}

	_, err = h.client.DeleteProduct(r.Context(), &pb.DeleteProductRequest{Id: id})
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

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid id:\n%v", err), http.StatusBadRequest)
		return
	}

	res, err := h.client.GetProduct(r.Context(), &pb.GetProductRequest{Id: id})
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

func (h *ProductHandler) ListProductsByCategory(w http.ResponseWriter, r *http.Request) {
	categoryIDStr := r.PathValue("category_id")
	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid category_id:\n%v", err), http.StatusBadRequest)
		return
	}

	res, err := h.client.ListProductsByCategory(r.Context(), &pb.ListProductsByCategoryRequest{CategoryId: categoryID})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.InvalidArgument:
				http.Error(w, st.Message(), http.StatusBadRequest)
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
