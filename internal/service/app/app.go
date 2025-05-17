package app

import (
	"context"

	"mathbot/internal/models"
)

type ProblemBook interface {
	Problem(ctx context.Context, id string) (models.Problem, error)
}

type App struct {
	pb ProblemBook
}

func New(pb ProblemBook) *App {
	return &App{pb: pb}
}
