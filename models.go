package main

import (
	"time"

	"github.com/google/uuid"
)

type Word struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Word      string    `json:"word"`
}


type Definition struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Definition string `json:"definition"`
	PartOfSpeech string `json:"part_of_speech"`
	WordID uuid.UUID `json:"word_id"`
}