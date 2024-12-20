package services

import (
	"auth/internal/models"
	"auth/internal/repositories"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user *models.User) error {
	// Validate user

	// Check if email exists
	if _, err := s.userRepo.FindByEmail(user.Email); err == nil {
		return fmt.Errorf("email already exists")
	}

	// Check if username exists
	if _, err := s.userRepo.FindByUsername(user.Username); err == nil {
		return fmt.Errorf("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)

	// Create user
	return s.userRepo.Create(user)
}
