package handler

import (
	"booking_chi_text/internal/domain/request"
	"booking_chi_text/internal/domain/response"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) initAuthGroup(router *chi.Mux) {
	router.Post("/register", h.CreateUser)
	router.Post("/login", h.Login)
	router.Post("/refresh-token", h.RefreshToken)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user request.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	err := h.services.Auth.Register(user.Name, user.Email, user.Password)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Register successfully", nil, http.StatusOK)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var user request.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	userData, err := h.services.Auth.Login(user.Email, user.Password)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	response.ResponseMessage(w, "Login Successfully", userData, http.StatusOK)
}

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var token struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	newAccessToken, err := h.services.Auth.RefreshToken(token.RefreshToken)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	response.ResponseMessage(w, "Refresh Token successfully", map[string]string{"access_token": newAccessToken}, http.StatusOK)
}
