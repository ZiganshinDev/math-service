package server

import (
	"net/http"

	"mathbot/internal/service/app"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func errorMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginErr := range c.Errors {
			if code, ok := app.Is4xxError(ginErr.Err); ok {
				c.JSON(code, gin.H{"error": ginErr.Err.Error()})
				return
			}

			logger.Error(ginErr.Err.Error())
			c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
	}
}

func metricsMiddleware(metrics Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		metrics.MetricMethod(c.Request)
	}
}

func tracerMiddleware(tracer Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer.Span(c.Request)
	}
}
