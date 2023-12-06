package repository

type TextBookingRepository interface {
	ReadAllBookings() []string
	CreateBookingWithDetails(tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda string)
	UpdateBookingWithDetails(bookingID int, tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda string)
	DeleteBooking(bookingID int)
}

type TextBookRepository interface {
	CreateBook(book_info string)
	ReadAllBooks() []string
	CreateBookWithDetails(title string, author string, publish_year string, isbn string)
	UpdateBookWithDetails(book_id int, title string, author string, publish_year string, isbn string)
	DeleteBook(book_id int)
}

type TextUserRepository interface {
	CreateUser(userInfo string)
	ReadAllUsers() []string
	FindByEmail(email string) map[string]string
	CreateUserWithDetails(name, email, password string)
	UpdateUserWithDetails(userID int, name, email, password string)
	DeleteUser(userID int)
}
