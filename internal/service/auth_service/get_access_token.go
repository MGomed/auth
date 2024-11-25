package authservice

import (
	"errors"
	"time"

	service_model "github.com/MGomed/auth/internal/model"
	token "github.com/MGomed/auth/pkg/token"
	jwt "github.com/dgrijalva/jwt-go"
)

func (s *service) GetAccessToken(
	refreshToken string,
	refreshSecretKey, accessSecretKey []byte,
	duration time.Duration,
) (string, error) {
	claims, err := token.VerifyToken(refreshToken, refreshSecretKey, &service_model.UserClaims{})
	if err != nil {
		return "", err
	}

	userClaims, ok := claims.(service_model.UserClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	accessToken, err := token.GenerateToken(&service_model.UserClaims{
		Email: userClaims.Email,
		Role:  userClaims.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}, accessSecretKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
