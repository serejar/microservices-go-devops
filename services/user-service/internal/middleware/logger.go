package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Logger returns a middleware that logs requests
func Logger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		latency := time.Since(start)

		// Prepare log fields
		fields := logrus.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       path,
			"latency":    latency,
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}

		if raw != "" {
			fields["query"] = raw
		}

		// Log request
		msg := "Request processed"
		statusCode := c.Writer.Status()

		switch {
		case statusCode >= 500:
			logger.WithFields(fields).Error(msg)
		case statusCode >= 400:
			logger.WithFields(fields).Warn(msg)
		default:
			logger.WithFields(fields).Info(msg)
		}
	}
}