package mathmaker

import (
	"context"
	"encoding/json"
	"fmt"

	"mathbot/internal/models"

	"github.com/google/uuid"
)

type Problems struct {
	Problems []Problem `json:"problems"`
}
type Problem struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Condition string    `json:"text"`
	Solution  string    `json:"solution"`
	Answer    string    `json:"answer"`
	// Optional
	// Files         []any  `json:"files"`
	// SolutionFiles []any  `json:"solution_files"`
	// Directory     string `json:"directory"`
	// Path          []struct {
	// 	ID   string `json:"id"`
	// 	Name string `json:"name"`
	// } `json:"path"`
}

func (m *Mathmaker) Problem(ctx context.Context, id uuid.UUID) (models.Problem, error) {
	const op = "mathmaker.problems.Problem"

	url := fmt.Sprintf("%s/api/problems/%s", m.baseURL, id)

	body, err := m.Get(ctx, url)
	if err != nil {
		return models.Problem{}, fmt.Errorf("%s: %w", op, err)
	}

	var mathProblem Problem

	if err := json.Unmarshal(body, &mathProblem); err != nil {
		return models.Problem{}, fmt.Errorf("%s: %w", op, err)
	}

	return mathProblemToModel(mathProblem), nil
}

func (m *Mathmaker) Problems(ctx context.Context) ([]models.Problem, error) {
	const op = "mathmaker.problems.Problem"

	url := fmt.Sprintf("%s/api/tree/default/search?page=1&with_answer=1", m.baseURL)

	body, err := m.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var mathProblems Problems

	if err := json.Unmarshal(body, &mathProblems); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	problems := make([]models.Problem, 0, len(mathProblems.Problems))
	for _, mathProblem := range mathProblems.Problems {
		problems = append(problems, mathProblemToModel(mathProblem))
	}

	return problems, nil
}

func mathProblemToModel(problem Problem) models.Problem {
	return models.Problem{
		ID:        problem.ID,
		Title:     problem.Title,
		Condition: problem.Condition,
		Solution:  problem.Solution,
		Answer:    problem.Answer,
	}
}
