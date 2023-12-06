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

func (h *Handler) initBookGroup(router *chi.Mux) {
	router.Route("/books", func(r chi.Router) {
		r.Use(middleware.TokenAuthMiddleware)

		r.Get("/", h.ReadAllBooks)
		r.Post("/create", h.CreateBook)
		r.Put("/update/{book_id:[0-9]+}", h.UpdateBook)
		r.Delete("/delete/{book_id:[0-9]+}", h.DeleteBook)
	})
}

func (h *Handler) ReadAllBooks(w http.ResponseWriter, r *http.Request) {
	books := h.services.Book.GetAllBooks()

	response.ResponseMessage(w, "Books already use data", books, http.StatusOK)
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book request.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	h.services.Book.CreateBook(book.Title, book.Author, book.PublishYear, book.ISBN)

	response.ResponseMessage(w, "Book created successfully", nil, http.StatusOK)
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookID, _ := strconv.Atoi(chi.URLParam(r, "book_id"))

	var book request.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.services.Book.UpdateBook(bookID, book.Title, book.Author, book.PublishYear, book.ISBN)

	response.ResponseMessage(w, fmt.Sprintf("Book with ID %d updated successfully", bookID), nil, http.StatusOK)
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID, _ := strconv.Atoi(chi.URLParam(r, "book_id"))

	h.services.Book.DeleteBook(bookID)

	response.ResponseMessage(w, fmt.Sprintf("Book with ID %d deleted successfully", bookID), nil, http.StatusOK)
}
