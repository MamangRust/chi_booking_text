package repository

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type textBooksRepository struct {
	FilePath string
}

func NewTextBookRepository(filepath string) *textBooksRepository {
	return &textBooksRepository{
		FilePath: filepath,
	}
}

func (t *textBooksRepository) CreateBook(book_info string) {
	file, err := os.OpenFile(t.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, book_info)
	if err != nil {
		fmt.Println(err)
	}
}

func (tbr *textBooksRepository) ReadAllBooks() []string {
	file, err := os.Open(tbr.FilePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	var books []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		books = append(books, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil
	}

	return books
}

func (tbr *textBooksRepository) CreateBookWithDetails(title string, author string, publish_year string, isbn string) {
	book_info := fmt.Sprintf("Title: %s, Author: %s, PublishYear: %s, ISBN: %s", title, author, publish_year, isbn)
	tbr.CreateBook(book_info)
}

func (tbr *textBooksRepository) UpdateBookWithDetails(book_id int, title string, author string, publish_year string, isbn string) {
	books := tbr.ReadAllBooks()
	if 0 <= book_id && book_id < len(books) {
		book_info := strings.Split(strings.TrimSpace(books[book_id]), ", ")

		// Prepare updated information based on provided or existing data
		updated_info := map[string]string{
			"Title":       title,
			"Author":      author,
			"PublishYear": publish_year,
			"ISBN":        isbn,
		}
		for key, value := range updated_info {
			if value == "" {
				updated_info[key] = strings.Split(book_info[getIndex(key)], ": ")[1]
			}
		}

		// Construct updated book information
		updated_book_info := fmt.Sprintf("Title: %s, Author: %s, PublishYear: %s, ISBN: %s\n", updated_info["Title"], updated_info["Author"], updated_info["PublishYear"], updated_info["ISBN"])

		// Update the book entry in the file
		books[book_id] = updated_book_info
		file, err := os.OpenFile(tbr.FilePath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		for _, book := range books {
			fmt.Fprintln(file, book)
		}
	} else {
		fmt.Println("Book ID out of range")
	}
}

func (t *textBooksRepository) DeleteBook(book_id int) {
	books := t.ReadAllBooks()

	if 0 <= book_id && book_id < len(books) {
		books = append(books[:book_id], books[book_id+1:]...)

		file, err := os.OpenFile(t.FilePath, os.O_WRONLY|os.O_TRUNC, 0644)

		if err != nil {
			fmt.Println(err)

			return
		}

		defer file.Close()

		for _, book := range books {
			fmt.Fprintln(file, book)
		}
	} else {
		fmt.Println("Book ID out of range")
	}
}

func getIndex(key string) int {
	switch key {
	case "Title":
		return 0
	case "Author":
		return 1
	case "PublishYear":
		return 2
	case "ISBN":
		return 3
	default:
		return -1
	}
}
