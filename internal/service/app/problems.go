package app

import (
	"context"
	"errors"
	"fmt"

	"mathbot/internal/models"
	"mathbot/internal/svcerr"

	"github.com/google/uuid"
)

var ErrProblemNotFound = errors.New("problem not found")

func (a *App) Problem(ctx context.Context, id uuid.UUID) (models.Problem, error) {
	const op = "app.problems.Problem"

	// problem, err := a.pb.Problem(ctx, id)
	// if err != nil {
	// 	return models.Problem{}, fmt.Errorf("%s: %w", op, err)
	// }

	problem := models.Problem{}

	if problem.ID == uuid.Nil {
		return models.Problem{}, fmt.Errorf("%s: %w", op, svcerr.New(svcerr.ErrNotFound, ErrProblemNotFound,
			fmt.Sprintf("id: %s", id)))
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
