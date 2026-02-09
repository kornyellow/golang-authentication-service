package service_test

import (
	"errors"
	"testing"

	"backend/internal/domain"
	"backend/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type mockUserRepo struct {
	mock.Mock
}

// Mock Create function
func (m *mockUserRepo) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Mock FindByUsername function
func (m *mockUserRepo) FindByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

// Test cases
func TestRegister_Success(t *testing.T) {
	mockRepo := new(mockUserRepo)
	authService := service.NewAuthService(mockRepo)

	req := domain.RegisterRequest{
		Username: "newuser",
		Password: "password123",
	}

	// 1. Searching for username returns not found (error)
	mockRepo.On("FindByUsername", "newuser").Return(nil, errors.New("record not found"))
	// 2. Creating user succeeds
	mockRepo.On("Create", mock.Anything).Return(nil)

	// Action
	err := authService.Register(req)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRegister_DuplicateUser(t *testing.T) {
	mockRepo := new(mockUserRepo)
	authService := service.NewAuthService(mockRepo)

	req := domain.RegisterRequest{
		Username: "existinguser",
		Password: "password123",
	}

	// 1. If find returns a user, it means username is taken
	foundUser := &domain.User{
		Username: "existinguser",
	}
	mockRepo.On("FindByUsername", "existinguser").Return(foundUser, nil)

	// Action
	err := authService.Register(req)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "username already exists", err.Error())
}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(mockUserRepo)
	authService := service.NewAuthService(mockRepo)

	password := "secret123"
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// 1. Try login with correct password
	mockUser := &domain.User{
		ID:           1,
		Username:     "korn",
		PasswordHash: string(hashed),
	}
	mockRepo.On("FindByUsername", "korn").Return(mockUser, nil)

	// Action
	token, err := authService.Login(domain.LoginRequest{
		Username: "korn",
		Password: password,
	})

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestLogin_WrongPassword(t *testing.T) {
	mockRepo := new(mockUserRepo)
	authService := service.NewAuthService(mockRepo)

	password := "correct_password"
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// 1. Try login with wrong password
	mockUser := &domain.User{
		Username:     "korn",
		PasswordHash: string(hashed),
	}
	mockRepo.On("FindByUsername", "korn").Return(mockUser, nil)

	// Action
	token, err := authService.Login(domain.LoginRequest{
		Username: "korn",
		Password: "wrong_password",
	})

	// Assert
	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Equal(t, "invalid username or password", err.Error())
}
