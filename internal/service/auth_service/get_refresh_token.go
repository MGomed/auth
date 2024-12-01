package authservice

import (
	"time"

	service_model "github.com/MGomed/auth/internal/model"
	service_errors "github.com/MGomed/auth/internal/service/errors"
	token "github.com/MGomed/common/token"
	jwt "github.com/dgrijalva/jwt-go"
)

func (s *service) GetRefreshToken(refreshToken string, secretKey []byte, duration time.Duration) (string, error) {
	claims, err := token.VerifyToken(refreshToken, secretKey, &service_model.UserClaims{})
	if err != nil {
		return "", err
	}

	userClaims, ok := claims.(service_model.UserClaims)
	if !ok {
		return "", service_errors.ErrInvalidToken
	}

	newRefreshToken, err := token.GenerateToken(service_model.UserClaims{
		Email: userClaims.Email,
		Role:  userClaims.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}, secretKey)
	if err != nil {
		return "", err
	}

	return newRefreshToken, nil
}
