package transport

import (
	"net/http"

	"github.com/go-chi/render"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (req *LoginRequest) Bind(r *http.Request) error {
	return nil
}

type LoginResponse struct {
	JWT            string `json:"token"`
	HTTPStatusCode int    `json:"-"`
}

func (resp *LoginResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, resp.HTTPStatusCode)
	return nil
}
