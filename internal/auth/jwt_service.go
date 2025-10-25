package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService interface {
	GenerateAccessToken(userID uint) (string, error)
	ValidateAccessToken(tokenStr string) (jwt.MapClaims, error)
}

type jwtService struct {
	secret string
	ttl    time.Duration
}

func NewJWTService(secret string, ttl time.Duration) JWTService {
	return &jwtService{secret: secret, ttl: ttl}
}

func (s *jwtService) GenerateAccessToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(s.ttl).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

func GenerateRefreshToken(userID string) (string, error) {
	token := uuid.NewString()
	// save to DB with userID, expiry = 7 days
	return token, nil
}

func (s *jwtService) ValidateAccessToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims.(jwt.MapClaims), nil
}
