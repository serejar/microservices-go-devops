package config

import (
	"errors"
	"os"
)

// Config holds the application configuration
type Config struct {
	Port            string
	UserServiceURL  string
	ProductServiceURL string
	Environment     string
}

// Load loads the configuration from environment variables
func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	userServiceURL := os.Getenv("USER_SERVICE_URL")
	if userServiceURL == "" {
		return nil, errors.New("USER_SERVICE_URL environment variable is required")
	}

	productServiceURL := os.Getenv("PRODUCT_SERVICE_URL")
	if productServiceURL == "" {
		return nil, errors.New("PRODUCT_SERVICE_URL environment variable is required")
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "development" // Default environment
	}

	return &Config{
		Port:            port,
		UserServiceURL:  userServiceURL,
		ProductServiceURL: productServiceURL,
		Environment:     environment,
	}, nil
}