package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// UserHandler handles user service-related requests
type UserHandler struct {
	userServiceURL string
	logger         *logrus.Logger
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userServiceURL string, logger *logrus.Logger) *UserHandler {
	return &UserHandler{
		userServiceURL: userServiceURL,
		logger:         logger,
	}
}

// GetUsers gets all users
func (h *UserHandler) GetUsers(c *gin.Context) {
	resp, err := http.Get(h.userServiceURL + "/users")
	if err != nil {
		h.logger.WithError(err).Error("Failed to get users from user service")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get users",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		h.logger.WithField("status_code", resp.StatusCode).Error("User service returned non-OK status")
		c.JSON(resp.StatusCode, gin.H{
			"error": "User service error",
		})
		return
	}

	var users interface{}
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		h.logger.WithError(err).Error("Failed to decode user service response")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to decode response",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser gets a user by ID
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get(h.userServiceURL + "/users/" + id)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get user from user service")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		h.logger.WithField("status_code", resp.StatusCode).Error("User service returned non-OK status")
		c.JSON(resp.StatusCode, gin.H{
			"error": "User service error",
		})
		return
	}

	var user interface{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		h.logger.WithError(err).Error("Failed to decode user service response")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to decode response",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", h.userServiceURL+"/users", c.Request.Body)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create request to user service")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := client.Do(req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create user in user service")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		h.logger.WithField("status_code", resp.StatusCode).Error("User service returned non-OK status")
		c.JSON(resp.StatusCode, gin.H{
			"error": "User service error",
		})
		return
	}

	var user interface{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		h.logger.WithError(err).Error("Failed to decode user service response")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to decode response",
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}