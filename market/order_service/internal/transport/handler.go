package transport

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	orderrepo "marketapi/orders/internal/repositories/order_repo"
	"marketapi/orders/internal/transport/middleware"
	orderusecase "marketapi/orders/internal/usecase"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	us        *orderusecase.OrderUseCase
	jwtSecret []byte
}

func New(usecase *orderusecase.OrderUseCase, secret string) *Handler {
	return &Handler{us: usecase, jwtSecret: []byte(secret)}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.ListOrders)
	r.Post("/", h.CreateOrder)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.GetOrder)
		r.With(middleware.JWTAuthMiddleware(h.jwtSecret)).Put("/", h.UpdateOrder)
		r.With(middleware.JWTAuthMiddleware(h.jwtSecret)).Delete("/", h.DeleteOrder)
		r.With(middleware.JWTAuthMiddleware(h.jwtSecret)).Put("/status", h.UpdateStatus)
	})
}

func (h *Handler) ListOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Received request to list orders")

	orders, err := h.us.ListOrders()
	if err != nil {
		log.Printf("Failed to fetch orders from usecase: %v", err)
		http.Error(w, "Failed to fetch orders: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if orders == nil {
		log.Println("No orders found, returning empty list")
		orders = make([]orderrepo.Order, 0)
	} else {
		log.Printf("Successfully fetched %d orders", len(orders))
	}

	if err := json.NewEncoder(w).Encode(orders); err != nil {
		log.Printf("Failed to encode orders to JSON: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	log.Println("Orders list successfully returned")
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req orderusecase.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Failed to decode order body: %v", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Creating order: %+v", req)

	created, err := h.us.CreateOrder(req)
	if err != nil {
		log.Printf("CreateOrder failed: %v", err)
		http.Error(w, "Failed to create order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(created); err != nil {
		log.Printf("Failed to encode created order: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	id := uint(id64)

	order, err := h.us.GetOrder(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch order: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(order)
}

func (h *Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	id := uint(id64)

	var o orderrepo.Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		log.Printf("Failed to decode update body: %v", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := h.us.UpdateOrder(id, o)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to update order: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(updated)
}

func (h *Handler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	id := uint(id64)

	var body struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Status == "" {
		http.Error(w, "Invalid or missing status field", http.StatusBadRequest)
		return
	}

	updated, err := h.us.UpdateStatus(id, body.Status)
	if err != nil {
		http.Error(w, "Failed to update status: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updated)
}

func (h *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	id := uint(id64)

	if err := h.us.DeleteOrder(id); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete order: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
