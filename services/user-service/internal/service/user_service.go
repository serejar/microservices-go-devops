package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/yourusername/go-microservices/user-service/internal/models"
	"github.com/yourusername/go-microservices/user-service/internal/repository"
)

// UserService handles business logic for users
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// GetUsers gets all users
func (s *UserService) GetUsers() ([]models.User, error) {
	return s.repo.GetUsers()
}

// GetUser gets a user by ID
func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

// CreateUser creates a new user
func (s *UserService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	now := time.Now()
	user := &models.User{
		ID:        uuid.New().String(),
		Email:     req.Email,
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	
	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}
	
	return user, nil
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(id string, req *models.UpdateUserRequest) (*models.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	
	if req.Email != "" {
		user.Email = req.Email
	}
	
	if req.Name != "" {
		user.Name = req.Name
	}
	
	user.UpdatedAt = time.Now()
	
	if err := s.repo.UpdateUser(user); err != nil {
		return nil, err
	}
	
	return user, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}