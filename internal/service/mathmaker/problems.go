package mathmaker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"mathbot/internal/models"
)

type Problem struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Text          string `json:"text"`
	Solution      string `json:"solution"`
	Answer        string `json:"answer"`
	Files         []any  `json:"files"`
	SolutionFiles []any  `json:"solution_files"`
	Directory     string `json:"directory"`
	Path          []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"path"`
}

func (m *Mathmaker) Problem(ctx context.Context, id string) (models.Problem, error) {
	const op = "mathmaker.problems.Problem"

	url := fmt.Sprintf("%s/problem/%s", m.baseURL, id)
	m.logger.Debugf("url", url)

	resp, err := http.Get(url)
	if err != nil {
		return models.Problem{}, fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Problem{}, fmt.Errorf("%s: %w", op, err)
	}

	var problem Problem

	if err := json.Unmarshal(body, &problem); err != nil {
		return models.Problem{}, fmt.Errorf("%s: %w", op, err)
	}

	return problemToModel(problem), nil
}

func problemToModel(problem Problem) models.Problem {
	return models.Problem{
		ID:       problem.ID,
		Title:    problem.Title,
		Text:     problem.Text,
		Solution: problem.Solution,
		Answer:   problem.Answer,
	}
}
