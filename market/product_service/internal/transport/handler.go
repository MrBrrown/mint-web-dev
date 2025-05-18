package transport

import (
	"database/sql"
	"encoding/json"
	productrepo "marketapi/products/internal/repositories/product_repo"
	productusecase "marketapi/products/internal/usecase"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	us *productusecase.ProductUseCase
}

func New(usecase *productusecase.ProductUseCase) *Handler {
	return &Handler{us: usecase}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/products", func(r chi.Router) {
		r.Get("/", h.ListProducts)
		r.Post("/", h.CreateProduct)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.GetProduct)
			r.Put("/", h.UpdateProduct)
			r.Delete("/", h.DeleteProduct)
		})
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
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	created, err := h.us.CreateProduct(p)
	if err != nil {
		http.Error(w, "failed to create product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
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
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}
	id := uint(id64)

	var p productrepo.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := h.us.UpdateProduct(id, p)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "product not found", http.StatusNotFound)
		} else {
			http.Error(w, "failed to update product: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(updated)
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
