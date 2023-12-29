// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package sqlc

import (
	"time"

	"github.com/google/uuid"
)

type Definition struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Definition   string
	PartOfSpeech string
	WordID       uuid.UUID
}

type Word struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Word      string
}
