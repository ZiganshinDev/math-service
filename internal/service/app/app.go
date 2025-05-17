package app

import (
	"context"

	"mathbot/internal/models"

	"github.com/google/uuid"
)

type ProblemBook interface {
	Problem(ctx context.Context, id uuid.UUID) (models.Problem, error)
	Problems(ctx context.Context) ([]models.Problem, error)
}

type App struct {
	pb ProblemBook
}

func New(pb ProblemBook) *App {
	return &App{pb: pb}
}
