package app

import (
	"context"
	"fmt"

	"mathbot/internal/models"

	"github.com/google/uuid"
)

func (a *App) Problem(ctx context.Context, id uuid.UUID) (models.Problem, error) {
	const op = "app.problems.Problem"

	problem, err := a.pb.Problem(ctx, id)
	if err != nil {
		return models.Problem{}, fmt.Errorf("%s: %w", op, err)
	}

	return problem, nil
}

func (a *App) Problems(ctx context.Context) ([]models.Problem, error) {
	const op = "app.problems.Problems"

	problems, err := a.pb.Problems(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return problems, nil
}
