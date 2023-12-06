package service

import (
	"booking_chi_text/internal/repository"
	"booking_chi_text/pkg/auth"
	"booking_chi_text/pkg/hash"
	"fmt"
	"time"
)

type textAuthService struct {
	Repository repository.TextUserRepository
	hash       hash.Hashing
	token      auth.TokenManager
}

func NewAuthService(repository repository.TextUserRepository, hash hash.Hashing, token auth.TokenManager) *textAuthService {
	return &textAuthService{
		Repository: repository,
		hash:       hash,
		token:      token,
	}
}

func (t *textAuthService) Register(name string, email string, password string) error {
	passwordHash, err := t.hash.HashPassword(password)
	if err != nil {
		return fmt.Errorf("error hashing password: %s", err.Error())
	}

	t.Repository.CreateUserWithDetails(name, email, passwordHash)

	return nil
}

func (t *textAuthService) Login(email string, password string) (map[string]string, error) {
	authEmail := t.Repository.FindByEmail(email)

	if authEmail == nil {
		return nil, fmt.Errorf("invalid Credentials")
	}

	passwordMatch := t.hash.ComparePassword(authEmail["Password"], password)
	if passwordMatch != nil {
		return nil, fmt.Errorf("incorrect password")
	}

	accessToken, err := t.token.CreateAccessToken(map[string]interface{}{"sub": authEmail["Email"]}, time.Hour) // Adjust the duration as needed
	if err != nil {
		return nil, fmt.Errorf("error generating access token: %s", err.Error())
	}

	refreshToken, err := t.token.CreateRefreshToken(map[string]interface{}{"sub": authEmail["Email"]})
	if err != nil {
		return nil, fmt.Errorf("error generating refresh token: %s", err.Error())
	}

	response := map[string]string{
		"name":          authEmail["Name"],
		"email":         authEmail["Email"],
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return response, nil
}

func (t *textAuthService) RefreshToken(token string) (string, error) {
	user, err := t.token.VerifyToken(token)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token")
	}

	access_token, err := t.token.CreateAccessToken(map[string]interface{}{"sub": user["Email"]}, time.Hour)
	if err != nil {
		return "", fmt.Errorf("error creating access token")
	}

	return access_token, nil
}
