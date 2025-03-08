package v1

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"math"
	"net/http"
	"time"
)

type Config struct {
	Host string
	Port string
}

type Handler struct {
	config *Config
	logger *zap.Logger
}

type HandlerParams struct {
	fx.In
	Config *Config
	Logger *zap.Logger
}

func NewHandler(params HandlerParams, router *gin.Engine) *Handler {
	handler := &Handler{
		config: params.Config,
		logger: params.Logger,
	}

	v1 := router.Group("/api/v1")
	v1.Use(LoggerMiddleware(params.Logger))
	v1.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	return handler
}

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()

		if len(c.Errors) > 0 {
			logger.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("[%s %d] %s (%dms)", c.Request.Method, statusCode, path, latency)
			if statusCode >= http.StatusInternalServerError {
				logger.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				logger.Warn(msg)
			} else {
				logger.Info(msg)
			}
		}
	}
}
