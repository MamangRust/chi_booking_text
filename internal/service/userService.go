package service

import (
	"booking_chi_text/internal/repository"
)

type textUsersService struct {
	UsersRepository repository.TextUserRepository
}

func NewTextUserService(repo repository.TextUserRepository) *textUsersService {
	return &textUsersService{
		UsersRepository: repo,
	}
}

func (s *textUsersService) FindByEmail(email string) map[string]string {
	return s.UsersRepository.FindByEmail(email)
}

func (s *textUsersService) CreateUser(name, email, password string) {
	s.UsersRepository.CreateUserWithDetails(name, email, password)
}

func (s *textUsersService) GetAllUsers() []string {
	return s.UsersRepository.ReadAllUsers()
}

func (s *textUsersService) UpdateUser(userID int, name, email, password string) {
	s.UsersRepository.UpdateUserWithDetails(userID, name, email, password)
}

func (s *textUsersService) DeleteUser(userID int) {
	s.UsersRepository.DeleteUser(userID)
}
