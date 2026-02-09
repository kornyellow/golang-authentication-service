package service

import (
	"errors"

	"backend/internal/domain"
	"backend/internal/repository"
	"backend/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req domain.RegisterRequest) error
	Login(req domain.LoginRequest) (string, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(req domain.RegisterRequest) error {
	// 1. Check for existing username
	_, err := s.repo.FindByUsername(req.Username)
	if err == nil {
		return errors.New("username already exists")
	}

	// 2. Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 3. Create User Object
	newUser := domain.User{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
	}

	// 4. Save to DB
	return s.repo.Create(&newUser)
}

func (s *authService) Login(req domain.LoginRequest) (string, error) {
	// 1. Find User by Username
	user, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// 2. Check Password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// 3. Generate JWT Token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
