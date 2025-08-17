package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CartHandler struct {
	client pb.CartServiceClient
}

func NewCartHandler(client pb.CartServiceClient) *CartHandler {
	return &CartHandler{client: client}
}

// AddToCart
func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if userID == 0 || !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req struct {
		ProductOfferID int64 `json:"product_offer_id"`
		Quantity       int32 `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request:\n%v", err), http.StatusBadRequest)
		return
	}

	res, err := h.client.AddToCart(r.Context(), &pb.AddToCartRequest{
		UserId:         userID,
		ProductOfferId: req.ProductOfferID,
		Quantity:       req.Quantity,
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

// UpdateCartItem
func (h *CartHandler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if userID == 0 || !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	offerIDStr := r.PathValue("offer_id")
	offerID, err := strconv.ParseInt(offerIDStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid offer_id:\n%v", err), http.StatusBadRequest)
		return
	}

	var req struct {
		Quantity int32 `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request:\n%v", err), http.StatusBadRequest)
		return
	}

	res, err := h.client.UpdateCartItem(r.Context(), &pb.UpdateCartItemRequest{
		UserId:         userID,
		ProductOfferId: offerID,
		Quantity:       req.Quantity,
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

// RemoveCartItem
func (h *CartHandler) RemoveCartItem(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if userID == 0 || !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	productOfferIDStr := r.PathValue("offer_id")
	productOfferID, err := strconv.ParseInt(productOfferIDStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid offer_id:\n%v", err), http.StatusBadRequest)
		return
	}

	_, err = h.client.RemoveCartItem(r.Context(), &pb.RemoveCartItemRequest{
		UserId:         userID,
		ProductOfferId: productOfferID,
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

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if userID == 0 || !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	res, err := h.client.GetCart(r.Context(), &pb.GetCartRequest{UserId: userID})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
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

func (h *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if userID == 0 || !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	_, err := h.client.ClearCart(r.Context(), &pb.ClearCartRequest{UserId: userID})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
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
