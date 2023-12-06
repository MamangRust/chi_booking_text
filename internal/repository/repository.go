package repository

type Repositories struct {
	User    TextUserRepository
	Book    TextBookRepository
	Booking TextBookingRepository
}

func NewRepository() *Repositories {
	return &Repositories{
		User:    NewTextUsersRepository("users.txt"),
		Book:    NewTextBookRepository("book.txt"),
		Booking: NewTextBookingRepository("booking.txt"),
	}
}
