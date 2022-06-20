package auth

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(userId int64) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

var SECRET_KEY = []byte("BWASTARTUP_keY")

// ValidateToken implements Service
func (s *service) ValidateToken(token string) (*jwt.Token, error) {

	tokenValid, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRET_KEY), nil

	})

	if err != nil {
		return nil, err
	}

	return tokenValid, nil
}

func (s *service) GenerateToken(userId int64) (string, error) {

	claim := jwt.MapClaims{
		"user_id": userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
