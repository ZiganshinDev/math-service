package internalhandler

import (
	"mathbot/internal/models"

	"github.com/google/uuid"
)

type Problems struct {
	Problems []Problem `json:"problems"`
}

type Problem struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Condition string    `json:"condition"`
	Answer    string    `json:"correct_answer"`
	Solution  string    `json:"solution"`
	// Optional
	// Methodics     string `json:"methodics"`
	// ConditionDesign string `json:"condition_design"`
	// AnswerOptions   string `json:"answer_options"`
	// Crucial bool `json:"crucial"`
	// SolutionDesign  string `json:"solution_design"`
	// AnswerDesign    string `json:"answer_design"`
	// Activity        []string      `json:"activity"`
	// Parent          []interface{} `json:"parent"`
	// Published       bool          `json:"published"`
	// Author          int           `json:"author"`
	// DateModified    time.Time     `json:"date_modified"`
	// Hints           []struct {
	// 	ID           string    `json:"id"`
	// 	Content      string    `json:"content"`
	// 	DateModified time.Time `json:"date_modified"`
	// } `json:"hints"`
}

func problemModelToResp(problem models.Problem) Problem {
	return Problem{
		ID:        problem.ID,
		Title:     problem.Title,
		Condition: problem.Condition,
		Answer:    problem.Answer,
		Solution:  problem.Solution,
	}
}
