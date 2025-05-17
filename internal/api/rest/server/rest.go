package server

import (
	"context"
	"fmt"
	"net/http"

	"mathbot/internal/api/rest/internalhandler"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type Metrics interface {
	MetricMethod(r *http.Request)
}

type Tracer interface {
	Span(r *http.Request)
}

type Server struct {
	srv *http.Server
}

func New(port string, app internalhandler.App, logger *zap.SugaredLogger) *Server {
	r := gin.Default()
	r.Use(errorMiddleware(logger))
	// r.Use(metricsMiddleware(metrics))
	// r.Use(tracerMiddleware(tracer))

	r.Handle("GET", "/metrics", func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})

	// handlers
	iHandler := internalhandler.New(app, logger)

	// grouping
	v1 := r.Group("api/v1")

	initHandles(v1, iHandler.Urls())

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r.Handler(),
	}

	return &Server{srv}
}

func initHandles(group *gin.RouterGroup, urls map[string]map[string]gin.HandlerFunc) {
	for path, handles := range urls {
		for method, handle := range handles {
			group.Handle(method, path, handle)
		}
	}
}

func (s *Server) Start() error {
	const op = "server.rest.Start"

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	const op = "server.rest.Shutdown"

	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
