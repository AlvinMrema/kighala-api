package main

import (
	"time"

	"github.com/AlvinMrema/kighala-api/internal/database"
	"github.com/google/uuid"
)

type Word struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Word      string    `json:"word"`
}

func databaseWordToWord(dbWord database.Word) Word {
	return Word{
		ID: dbWord.ID,
		CreatedAt: dbWord.CreatedAt,
		UpdatedAt: dbWord.UpdatedAt,
		Word: dbWord.Word,
	}
}

func databaseWordsToWords(dbWords []database.Word) []Word {
	words := []Word{}
	for _, dbWord := range dbWords {
		words = append(words, databaseWordToWord(dbWord))
	}
	return words
}

type Definition struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Definition   string    `json:"definition"`
	PartOfSpeech string    `json:"part_of_speech"`
	WordID       uuid.UUID `json:"word_id"`
}
