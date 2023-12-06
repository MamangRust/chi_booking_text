package repository

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type textBookingRepository struct {
	FilePath string
}

func NewTextBookingRepository(filepath string) *textBookingRepository {
	return &textBookingRepository{FilePath: filepath}
}

func (t *textBookingRepository) ReadAllBookings() []string {
	data, err := os.ReadFile(t.FilePath)

	if err != nil {
		fmt.Println(err)

		return []string{}
	}

	return strings.Split(string(data), "\n")
}

func (t *textBookingRepository) CreateBooking(bookingInfo string) {
	data := []byte(bookingInfo + "\n")

	err := os.WriteFile(t.FilePath, data, 0644)

	if err != nil {
		fmt.Println(err)
	}
}

func (t *textBookingRepository) CreateBookingWithDetails(tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda string) {
	bookingInfo := fmt.Sprintf("TglPinjam: %s, UserID: %s, TglKembali: %s, TglPengembalian: %s, Status: %s", tglPinjam, userID, tglKembali, tglPengembalian, status)

	layout := "2006-01-02"
	tglKembaliTime, _ := time.Parse(layout, tglKembali)
	tglPengembalianTime, _ := time.Parse(layout, tglPengembalian)

	if tglPengembalianTime.After(tglKembaliTime) {
		timeDifference := tglPengembalianTime.Sub(tglKembaliTime)
		hoursLate := int(timeDifference.Hours())
		denda := hoursLate * 10000
		bookingInfo += fmt.Sprintf(", Denda: %d", denda)
	}

	t.CreateBooking(bookingInfo)
}

func (t *textBookingRepository) UpdateBookingWithDetails(bookingID int, tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda string) {
	bookings := t.ReadAllBookings()

	if bookingID >= 0 && bookingID < len(bookings) {
		booking := strings.Split(bookings[bookingID], ", ")

		updatedTglPinjam := tglPinjam
		updatedUserID := userID
		updatedTglKembali := tglKembali
		updatedTglPengembalian := tglPengembalian
		updatedStatus := status
		updatedTotalDenda := totalDenda

		if userID == "" {
			updatedUserID = strings.Split(booking[1], ": ")[1]
		}
		if tglKembali == "" {
			updatedTglKembali = strings.Split(booking[2], ": ")[1]
		}
		if tglPengembalian == "" {
			updatedTglPengembalian = strings.Split(booking[3], ": ")[1]
		}
		if status == "" {
			updatedStatus = strings.Split(booking[4], ": ")[1]
		}
		if totalDenda == "" {
			updatedTotalDenda = strings.Split(booking[5], ": ")[1]
		}

		if tglPengembalian != "" && tglKembali != "" {
			layout := "2006-01-02"

			tglKembaliTime, _ := time.Parse(layout, updatedTglKembali)
			tglPengembalianTime, _ := time.Parse(layout, updatedTglPengembalian)

			if tglPengembalianTime.After(tglKembaliTime) {
				timeDifference := tglPengembalianTime.Sub(tglKembaliTime)
				hourLate := int(timeDifference.Hours())
				denda := hourLate * 10000
				updatedTotalDendaInt := 0

				if updatedTotalDenda != "" {
					updatedTotalDendaInt, _ = strconv.Atoi(updatedTotalDenda)
				}

				updatedTotalDenda = strconv.Itoa(updatedTotalDendaInt + denda)
				updatedStatus = "Late Return"
			}
		}

		updatedBookingInfo := fmt.Sprintf("TglPinjam: %s, UserID: %s, TglKembali: %s, TglPengembalian: %s, Status: %s, TotalDenda: %s\n", updatedTglPinjam, updatedUserID, updatedTglKembali, updatedTglPengembalian, updatedStatus, updatedTotalDenda)
		bookings[bookingID] = updatedBookingInfo

		updatedData := []byte(strings.Join(bookings, "\n"))

		err := os.WriteFile(t.FilePath, updatedData, 0644)

		if err != nil {
			fmt.Println(err)
		}

	} else {
		fmt.Println("Booking ID out of range")
	}
}

func (t *textBookingRepository) DeleteBooking(bookingID int) {
	bookings := t.ReadAllBookings()

	if bookingID >= 0 && bookingID < len(bookings) {
		bookings = append(bookings[:bookingID], bookings[bookingID+1:]...)
		updatedData := []byte(strings.Join(bookings, "\n"))
		err := os.WriteFile(t.FilePath, updatedData, 0644)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Booking ID Out of range")
	}
}
