package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

func InitLogger() (*zap.SugaredLogger, error) {
	var logger *zap.Logger

	if os.Getenv("GIN_MODE") == "release" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	defer func() {
		if err := logger.Sync(); err != nil {
			// Handle the error in some manner, e.g., log it or ignore it based on context
		}
	}() // flushes buffer, if any
	return logger.Sugar(), nil
}

func GinLoggerHandler(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		logger.Infow("request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", duration,
			"client_ip", c.ClientIP(),
		)
	}
}
