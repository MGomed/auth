package token

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	errMethodUnexpected = errors.New("unexpected token signing method")
)

// GenerateToken generates jwt-token
func GenerateToken(claims jwt.Claims, secretKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

// VerifyToken verifies jwt-token
func VerifyToken(tokenStr string, secretKey []byte, claims jwt.Claims) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errMethodUnexpected
		}

		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims, nil
}
