package internalhandler

import (
	"context"

	"mathbot/internal/models"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type App interface {
	Problem(ctx context.Context, id uuid.UUID) (models.Problem, error)
	Problems(ctx context.Context) ([]models.Problem, error)
}

type Handler struct {
	app    App
	logger *zap.SugaredLogger
}

func New(app App, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		app:    app,
		logger: logger,
	}
}
