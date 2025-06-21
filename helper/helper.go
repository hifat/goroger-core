package helper

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	core "github.com/hifat/goroger-core"
	"github.com/jinzhu/copier"
)

type helper struct{}

func New() core.Helper {
	return &helper{}
}

func (h *helper) Copy(toValue any, fromValue any) error {
	return copier.Copy(toValue, fromValue)
}

var ErrExpectedSigningMethod = errors.New("expected signing method")
var ErrInvalidToken = errors.New("invalid token")

func (h *helper) ParseToken(secret string, tokenString string, claims jwt.Claims) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrExpectedSigningMethod
		}

		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims, ok := token.Claims.(jwt.Claims); ok && token.Valid {
		return tokenClaims, nil
	}

	return nil, ErrInvalidToken
}
