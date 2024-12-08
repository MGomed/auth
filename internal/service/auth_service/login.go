package authservice

import (
	"context"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	bcrypt "golang.org/x/crypto/bcrypt"

	service_model "github.com/MGomed/auth/internal/model"
	token "github.com/MGomed/common/token"
)

// Login generates refreshToken and accessToken for authorized user
func (s *service) Login(
	ctx context.Context,
	email, password string,
	refreshSecretKey, accessSecretKey []byte,
	refreshExpTime, accessExpTime time.Duration,
) (string, string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		s.log.Error("Failed to get user by email: %v", err)

		return "", "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
		s.log.Error("Failed user password matching!")

		return "", "", err
	}

	refreshToken, err := token.GenerateToken(service_model.UserClaims{
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshExpTime).Unix(),
		},
	}, refreshSecretKey)
	if err != nil {
		s.log.Error("Failed to generate refresh token: %v", err)

		return "", "", err
	}

	accessToken, err := token.GenerateToken(service_model.UserClaims{
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessExpTime).Unix(),
		},
	}, accessSecretKey)
	if err != nil {
		s.log.Error("Failed to generate access token: %v", err)

		return "", "", err
	}

	s.log.Debug("Successfully generate refresh (%v) and access (%v) tokens", refreshToken, accessToken)

	return refreshToken, accessToken, nil
}
