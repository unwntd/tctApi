package auth

import (
	"errors"
	"tctApi/internal/user/repository"
	"time"
)

type AuthService interface {
	Login(email, password string) (*TokenPair, string, error)
	Refresh(userID uint, refreshID string) (*TokenPair, string, error)
}

type authService struct {
	jwtService  JWTService
	userRepo    repository.UserRepository
	refreshRepo RefreshTokenRepository
}

func NewAuthService(jwt JWTService, userRepo repository.UserRepository, refreshRepo RefreshTokenRepository) AuthService {
	return &authService{jwtService: jwt, userRepo: userRepo, refreshRepo: refreshRepo}
}

func (s *authService) Login(email, password string) (*TokenPair, string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil || !user.CheckPassword(password) {
		return nil, "", errors.New("invalid credentials")
	}

	accessToken, err := s.jwtService.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	refresh, err := s.refreshRepo.Create(user.ID, time.Now().Add(7*24*time.Hour))
	if err != nil {
		return nil, "", err
	}

	return &TokenPair{AccessToken: accessToken}, refresh.Token, nil
}

func (s *authService) Refresh(userID uint, refreshID string) (*TokenPair, string, error) {
	validToken, err := s.refreshRepo.Validate(refreshID, userID)
	if err != nil {
		return nil, "", errors.New("invalid or expired refresh token")
	}

	// rotate refresh token
	_ = s.refreshRepo.Revoke(validToken.Token, validToken.UserID)

	newAccess, _ := s.jwtService.GenerateAccessToken(userID)
	newRefresh, err := s.refreshRepo.Create(userID, time.Now().Add(7*24*time.Hour))
	if err != nil {
		return nil, "", err
	}

	return &TokenPair{AccessToken: newAccess}, newRefresh.Token, nil
}
