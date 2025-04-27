package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/go-microservices/user-service/internal/models"
	"github.com/yourusername/go-microservices/user-service/internal/service"
)

// UserHandler handles user-related requests
type UserHandler struct {
	service *service.UserService
	logger  *logrus.Logger
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(service *service.UserService, logger *logrus.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

// GetUsers gets all users
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		h.logger.WithError(err).Error("Failed to get users")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get users",
		})
		return
	}
	
	c.JSON(http.StatusOK, users)
}

// GetUser gets a user by ID
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	
	user, err := h.service.GetUser(id)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get user")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	
	c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request payload")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}
	
	user, err := h.service.CreateUser(&req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create user")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	
	c.JSON(http.StatusCreated, user)
}

// UpdateUser updates a user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	
	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request payload")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}
	
	user, err := h.service.UpdateUser(id, &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to update user")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	
	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	
	if err := h.service.DeleteUser(id); err != nil {
		h.logger.WithError(err).Error("Failed to delete user")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	
	c.JSON(http.StatusNoContent, nil)
}