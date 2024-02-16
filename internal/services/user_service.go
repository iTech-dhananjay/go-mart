package services

import (
	"github.com/yourusername/yourproject/internal/models"
	"github.com/yourusername/yourproject/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(), // Initialize user repository
	}
}

func (us *UserService) GetAllUsers() ([]models.User, error) {
	return us.userRepo.GetAllUsers()
}
