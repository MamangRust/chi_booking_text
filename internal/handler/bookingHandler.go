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

func (h *Handler) initBookingGroup(router *chi.Mux) {
	router.Route("/bookings", func(r chi.Router) {
		r.Use(middleware.TokenAuthMiddleware)
		r.Get("/", h.ReadAllBookings)
		r.Post("/create", h.CreateBooking)
		r.Put("/update/{booking_id:[0-9]+}", h.UpdateBooking)
		r.Delete("/delete/{booking_id:[0-9]+}", h.DeleteBooking)
	})
}

func (h *Handler) ReadAllBookings(w http.ResponseWriter, r *http.Request) {
	bookings := h.services.Booking.GetAllBookings()

	response.ResponseMessage(w, "Bookings already in use", bookings, http.StatusOK)
}

func (h *Handler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking request.Booking

	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	h.services.Booking.CreateBooking(
		booking.TglPinjam,
		strconv.Itoa(booking.UserID),
		booking.TglKembali,
		booking.TglPengembalian,
		booking.Status,
		strconv.Itoa(booking.TotalDenda),
	)

	response.ResponseMessage(w, "Booking created successfully", nil, http.StatusOK)
}

func (h *Handler) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	bookingID, _ := strconv.Atoi(chi.URLParam(r, "booking_id"))

	var booking request.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.services.Booking.UpdateBooking(
		bookingID,
		booking.TglPinjam,
		strconv.Itoa(booking.UserID),
		booking.TglKembali,
		booking.TglPengembalian,
		booking.Status,
		strconv.Itoa(booking.TotalDenda),
	)

	response.ResponseMessage(w, fmt.Sprintf("Booking with ID %d updated successfully", bookingID), nil, http.StatusOK)
}

func (h *Handler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	bookingID, _ := strconv.Atoi(chi.URLParam(r, "booking_id"))

	h.services.Booking.DeleteBooking(bookingID)

	response.ResponseMessage(w, fmt.Sprintf("Booking with ID %d deleted successfully", bookingID), nil, http.StatusOK)
}
