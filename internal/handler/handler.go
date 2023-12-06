package handler

import (
	"booking_chi_text/internal/service"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	h.InitApi(router)

	return router
}

func (h *Handler) InitApi(router *chi.Mux) {
	h.initAuthGroup(router)
	h.initUserGroup(router)
	h.initBookGroup(router)
	h.initBookingGroup(router)
}
