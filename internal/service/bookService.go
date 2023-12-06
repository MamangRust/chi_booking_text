package service

import "booking_chi_text/internal/repository"

type textBookService struct {
	repository repository.TextBookRepository
}

func NewTextBookService(repository repository.TextBookRepository) *textBookService {
	return &textBookService{
		repository: repository,
	}
}

func (t *textBookService) GetAllBooks() []string {
	return t.repository.ReadAllBooks()
}

func (t *textBookService) CreateBook(title string, author string, publish_year string, isbn string) {
	t.repository.CreateBookWithDetails(title, author, publish_year, isbn)
}

func (t *textBookService) UpdateBook(book_id int, title string, author string, publish_year string, isbn string) {
	t.repository.UpdateBookWithDetails(book_id, title, author, publish_year, isbn)
}

func (t *textBookService) DeleteBook(book_id int) {
	t.repository.DeleteBook(book_id)
}
