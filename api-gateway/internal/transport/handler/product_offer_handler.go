package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alibekkenny/simple-marketplace/api-gateway/internal/middleware"
	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductOfferHandler struct {
	client pb.ProductOfferServiceClient
}

func NewProductOfferHandler(client pb.ProductOfferServiceClient) *ProductOfferHandler {
	return &ProductOfferHandler{client: client}
}

func (h *ProductOfferHandler) CreateProductOffer(w http.ResponseWriter, r *http.Request) {
	supplierID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if supplierID == 0 || !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	productIDStr := r.PathValue("product_id")
	productID, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		http.Error(w, "product_id is required", http.StatusBadRequest)
		return
	}

	var req struct {
		Stock    int32   `json:"stock"`
		Price    float64 `json:"price"`
		IsActive bool    `json:"is_active"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request:\n%v", err), http.StatusBadRequest)
		return
	}

	res, err := h.client.CreateProductOffer(r.Context(), &pb.CreateProductOfferRequest{
		Stock:      req.Stock,
		Price:      req.Price,
		IsActive:   req.IsActive,
		ProductId:  productID,
		SupplierId: supplierID,
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

func (h *ProductOfferHandler) UpdateProductOffer(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("offer_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid offer_id:\n%v", err), http.StatusBadRequest)
		return
	}

	var req struct {
		Stock    int32   `json:"stock"`
		Price    float64 `json:"price"`
		IsActive bool    `json:"is_active"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request:\n%v", err), http.StatusBadRequest)
		return
	}

	res, err := h.client.UpdateProductOffer(r.Context(), &pb.UpdateProductOfferRequest{
		Id:       id,
		Stock:    req.Stock,
		Price:    req.Price,
		IsActive: req.IsActive,
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

func (h *ProductOfferHandler) DeleteProductOffer(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("offer_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid offer_id:\n%v", err), http.StatusBadRequest)
		return
	}

	_, err = h.client.DeleteProductOffer(r.Context(), &pb.DeleteProductOfferRequest{
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

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductOfferHandler) GetProductOfferByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("offer_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid offer_id:\n%v", err), http.StatusBadRequest)
		return
	}

	res, err := h.client.GetProductOffer(r.Context(), &pb.GetProductOfferRequest{
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
	json.NewEncoder(w).Encode(res)
}

func (h *ProductOfferHandler) ListProductOffersByProductID(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.PathValue("product_id")
	productID, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid product_id:\n%v", err), http.StatusBadRequest)
		return
	}

	res, err := h.client.GetProductOffersByProduct(r.Context(), &pb.GetProductOffersByProductRequest{
		ProductId: productID,
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

func (h *ProductOfferHandler) ListProductOffersBySupplierID(w http.ResponseWriter, r *http.Request) {
	supplierID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if supplierID == 0 || !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	res, err := h.client.GetProductOffersBySupplier(r.Context(), &pb.GetProductOffersBySupplierRequest{
		SupplierId: supplierID,
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
