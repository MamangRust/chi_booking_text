package auth

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey                 = []byte("09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7")
	algorithm                 = jwt.SigningMethodHS256
	accessTokenExpireMinutes  = 3000
	refreshTokenExpireMinutes = 60 * 24 * 7 // 7 days
)

type TokenManager interface {
	CreateAccessToken(data map[string]interface{}, expiresDelta time.Duration) (string, error)
	CreateRefreshToken(data map[string]interface{}) (string, error)
	VerifyToken(tokenString string) (map[string]interface{}, error)
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{signingKey: signingKey}, nil
}

func (m *Manager) CreateAccessToken(data map[string]interface{}, expiresDelta time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(expiresDelta).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (m *Manager) CreateRefreshToken(data map[string]interface{}) (string, error) {
	return m.CreateAccessToken(data, time.Minute*time.Duration(refreshTokenExpireMinutes))
}

func (m *Manager) VerifyToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := fmt.Sprintf("%v", claims["sub"])

		users, err := ioutil.ReadFile("users.txt")
		if err != nil {
			return nil, err
		}

		for _, userLine := range strings.Split(string(users), "\n") {
			info := strings.Split(userLine, ", ")
			if len(info) == 3 && email == strings.Split(info[1], ": ")[1] {
				return map[string]interface{}{
					"name":     strings.Split(info[0], ": ")[1],
					"email":    strings.Split(info[1], ": ")[1],
					"password": strings.Split(info[2], ": ")[1],
				}, nil
			}
		}

		return nil, fmt.Errorf("user not found")
	}

	return nil, fmt.Errorf("invalid token")
}
