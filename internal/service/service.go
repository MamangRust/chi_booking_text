package service

import (
	"booking_chi_text/internal/repository"
	"booking_chi_text/pkg/auth"
	"booking_chi_text/pkg/hash"
)

type Service struct {
	User    TextUserService
	Book    TextBookService
	Booking TextBookingService
	Auth    textAuthService
}

type Deps struct {
	Repository *repository.Repositories
	Hashing    hash.Hashing
	Token      auth.TokenManager
}

func NewService(deps Deps) *Service {
	return &Service{
		User:    NewTextUserService(deps.Repository.User),
		Book:    NewTextBookService(deps.Repository.Book),
		Booking: NewTextBookingService(deps.Repository.Booking),
		Auth:    *NewAuthService(deps.Repository.User, deps.Hashing, deps.Token),
	}
}
