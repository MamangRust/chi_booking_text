package service

import "booking_chi_text/internal/repository"

type textBookingService struct {
	BookingRepository repository.TextBookingRepository
}

func NewTextBookingService(repo repository.TextBookingRepository) *textBookingService {
	return &textBookingService{BookingRepository: repo}
}

func (t *textBookingService) GetAllBookings() []string {
	return t.BookingRepository.ReadAllBookings()
}

func (t *textBookingService) CreateBooking(tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda string) {
	t.BookingRepository.CreateBookingWithDetails(tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda)
}

func (t *textBookingService) UpdateBooking(bookingID int, tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda string) {
	t.BookingRepository.UpdateBookingWithDetails(bookingID, tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda)
}

func (t *textBookingService) DeleteBooking(bookingID int) {
	t.BookingRepository.DeleteBooking(bookingID)
}
