package transport

import (
	"encoding/json"
	"log"
	"marketapi/auth/internal/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Handler struct {
	us *usecase.AuthUseCase
}

func NewAuthHandler(usecase *usecase.AuthUseCase) *Handler {
	return &Handler{us: usecase}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/login", h.loginFunc)
}

func (h *Handler) loginFunc(w http.ResponseWriter, r *http.Request) {
	req := LoginRequest{}
	err := render.Bind(r, &req)
	if err != nil {
		log.Printf("error : %s", err.Error())
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.us.TryLogin(req.Login, req.Password)
	if err != nil {
		log.Printf("error %s : %s", req.Login, err.Error())
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	resp := LoginResponse{JWT: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
