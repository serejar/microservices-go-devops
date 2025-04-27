package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/go-microservices/api-gateway/internal/config"
	"github.com/yourusername/go-microservices/api-gateway/internal/handlers"
	"github.com/yourusername/go-microservices/api-gateway/internal/middleware"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize router
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Logger(logger))
	router.Use(middleware.PrometheusMetrics())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
			"name":   "api-gateway",
		})
	})

	// Prometheus metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// API routes
	apiGroup := router.Group("/api")
	{
		// User service routes
		userHandler := handlers.NewUserHandler(cfg.UserServiceURL, logger)
		apiGroup.GET("/users", userHandler.GetUsers)
		apiGroup.GET("/users/:id", userHandler.GetUser)
		apiGroup.POST("/users", userHandler.CreateUser)

		// Product service routes
		productHandler := handlers.NewProductHandler(cfg.ProductServiceURL, logger)
		apiGroup.GET("/products", productHandler.GetProducts)
		apiGroup.GET("/products/:id", productHandler.GetProduct)
		apiGroup.POST("/products", productHandler.CreateProduct)
	}

	// Create HTTP server
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		logger.Infof("Starting API Gateway on port %s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Create context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Info("Server exiting")
}