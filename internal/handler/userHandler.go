package handler

import (
	"booking_chi_text/internal/domain/request"
	"booking_chi_text/internal/domain/response"
	"booking_chi_text/internal/middleware"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) initUserGroup(router *chi.Mux) {

	router.Route("/users", func(r chi.Router) {
		r.Use(middleware.TokenAuthMiddleware)
		r.Get("/", h.ReadAllUsers)
		r.Post("/create", h.CreateUser)
		r.Put("/update/{user_id:[0-9]+}", h.UpdateUser)
		r.Put("/delete/{user_id:[0-9]+}", h.DeleteUser)
	})
}

func (h *Handler) ReadAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.services.User.GetAllUsers()

	response.ResponseMessage(w, "User already use data", users, http.StatusOK)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user request.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	h.services.User.CreateUser(user.Name, user.Email, user.Password)

	response.ResponseMessage(w, "User created successfully", nil, http.StatusOK)

}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "user_id"))

	var user request.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.services.User.UpdateUser(userID, user.Name, user.Email, user.Password)

	response.ResponseMessage(w, fmt.Sprintf("User with ID %d updated successfully", userID), nil, http.StatusOK)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "user_id"))

	h.services.User.DeleteUser(userID)

	response.ResponseMessage(w, fmt.Sprintf("User with ID %d deleted successfully", userID), nil, http.StatusOK)

}
