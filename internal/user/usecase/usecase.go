package usecase

import (
	"errors"
	"tctApi/internal/user"
	"tctApi/internal/user/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(input *user.RegisterRequest) (*user.User, error)
	Login(input *user.LoginRequest) (string, *user.User, error)
}

type userUsecase struct {
	repo      repository.UserRepository
	jwtSecret string
}

func NewUserUsecase(repo repository.UserRepository, jwtSecret string) UserUsecase {
	return &userUsecase{repo: repo, jwtSecret: jwtSecret}
}

func (u *userUsecase) Register(input *user.RegisterRequest) (*user.User, error) {
	// Check if email already exists
	existing, _ := u.repo.FindByEmail(input.Email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	newUser := &user.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
	}

	if err := u.repo.Create(newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *userUsecase) Login(input *user.LoginRequest) (string, *user.User, error) {
	user, err := u.repo.FindByEmail(input.Email)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}
