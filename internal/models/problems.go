package models

import "github.com/google/uuid"

type Problem struct {
	ID        uuid.UUID
	Title     string
	Condition string
	Solution  string
	Answer    string
}
