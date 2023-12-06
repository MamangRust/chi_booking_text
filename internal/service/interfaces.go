package service

type TextUserService interface {
	FindByEmail(email string) map[string]string
	CreateUser(name, email, password string)
	GetAllUsers() []string
	UpdateUser(userID int, name, email, password string)
	DeleteUser(userID int)
}

type TextBookService interface {
	GetAllBooks() []string
	CreateBook(title string, author string, publish_year string, isbn string)
	UpdateBook(book_id int, title string, author string, publish_year string, isbn string)
	DeleteBook(book_id int)
}

type TextBookingService interface {
	GetAllBookings() []string
	CreateBooking(tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda string)
	UpdateBooking(bookingID int, tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda string)
	DeleteBooking(bookingID int)
}

type TextAuthService interface {
	Register(name string, email string, password string) error
	Login(email string, password string) (map[string]string, error)
	RefreshToken(token string) (string, error)
}
