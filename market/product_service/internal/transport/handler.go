package transport

import (
	"database/sql"
	"encoding/json"
	"log"
	productrepo "marketapi/products/internal/repositories/product_repo"
	"marketapi/products/internal/transport/middleware"
	productusecase "marketapi/products/internal/usecase"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	us        *productusecase.ProductUseCase
	jwtSecret []byte
}

func New(usecase *productusecase.ProductUseCase, secret string) *Handler {
	return &Handler{us: usecase, jwtSecret: []byte(secret)}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.ListProducts)
	r.Post("/batch", h.batchProduct)
	r.With(middleware.JWTAuthMiddleware(h.jwtSecret)).Post("/", h.CreateProduct)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.GetProduct)
		r.With(middleware.JWTAuthMiddleware(h.jwtSecret)).Put("/", h.UpdateProduct)
		r.With(middleware.JWTAuthMiddleware(h.jwtSecret)).Delete("/", h.DeleteProduct)
	})
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.us.ListProducts()
	if err != nil {
		http.Error(w, "failed to fetch products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p productrepo.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if p.Name == "" || p.Price <= 0 {
		log.Printf("Validation error: missing name or invalid price. name='%s', price=%f", p.Name, p.Price)
		http.Error(w, "Missing required fields: name and valid price are required", http.StatusBadRequest)
		return
	}

	log.Printf("Attempting to create product: %+v", p)

	created, err := h.us.CreateProduct(p)
	if err != nil {
		log.Printf("Error creating product in usecase: %v", err)
		http.Error(w, "Failed to create product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Product created successfully: %+v", created)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(created); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}
	id := uint(id64)

	p, err := h.us.GetProduct(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "product not found", http.StatusNotFound)
		} else {
			http.Error(w, "failed to fetch product: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(p)
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("Invalid product ID in URL: %v", err)
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	id := uint(id64)

	var p productrepo.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Attempting to update product with ID=%d: %+v", id, p)

	updated, err := h.us.UpdateProduct(id, p)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Product with ID=%d not found", id)
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			log.Printf("Failed to update product with ID=%d: %v", id, err)
			http.Error(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	log.Printf("Product updated successfully: %+v", updated)

	if err := json.NewEncoder(w).Encode(updated); err != nil {
		log.Printf("Failed to encode response for updated product ID=%d: %v", id, err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}
	id := uint(id64)

	if err := h.us.DeleteProduct(id); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "product not found", http.StatusNotFound)
		} else {
			http.Error(w, "failed to delete product: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) batchProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ids []uint
	if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if len(ids) == 0 {
		http.Error(w, "empty id list", http.StatusBadRequest)
		return
	}

	products := make([]productrepo.Product, 0)
	for _, id := range ids {
		p, err := h.us.GetProduct(id)
		if err != nil {
			continue
		}

		products = append(products, p)
	}

	json.NewEncoder(w).Encode(products)
}
