package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"tube/pkg/config"
	"tube/pkg/repository"
)

type AuthService interface {
	CreateToken(data map[string]interface{}) (string, error)
	ValidateToken(tokenString string) (map[string]interface{}, error)
}

type AuthServiceImpl struct {
	repo       repository.Repository
	hmacSecret []byte
	duration   int
}

func NewAuthService(repo repository.Repository, config *config.Application) AuthService {
	return &AuthServiceImpl{
		repo:       repo,
		hmacSecret: []byte(config.Auth.SecretKey),
		duration:   config.Auth.Duration,
	}
}

func (t *AuthServiceImpl) CreateToken(data map[string]interface{}) (string, error) {
	claims := jwt.MapClaims(data)
	claims["exp"] = time.Now().AddDate(0, 0, t.duration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(t.hmacSecret)
}

func (t *AuthServiceImpl) ValidateToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return t.hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token structure")
	}
	return claims, err
}
