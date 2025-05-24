package server

import (
	"errors"
	"net/http"

	"mathbot/internal/svcerr"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func errorMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginErr := range c.Errors {
			var svcErr *svcerr.SvcErr

			var statusCode int

			if !errors.As(ginErr.Err, &svcErr) {
				logger.Error(ginErr.Err.Error())
				c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				break
			}

			switch svcErr.BaseErr {
			case svcerr.ErrBadReq:
				statusCode = http.StatusBadRequest
			case svcerr.ErrNotFound:
				statusCode = http.StatusNotFound
			default:
				logger.Error(ginErr.Err.Error())
				c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				return
			}

			c.JSON(statusCode, gin.H{"error": ginErr.Err.Error()})
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
