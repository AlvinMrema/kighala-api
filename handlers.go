package main

import (
	"fmt"
	"time"

	"github.com/AlvinMrema/kighala-api/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateWord(c *fiber.Ctx) error {
	word := Word{}

	if err := c.BodyParser(&word); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Check if 'word' already exists
	value, _ := apiCfg.DB.GetWordByValue(c.Context(), word.Word)

	if value.Word == word.Word {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   fmt.Sprintf("Duplicate: '%s' already exists in Database.", value.Word),
		})
	}

	// Creating the 'word' to database
	db, err := apiCfg.DB.CreateWord(c.Context(), database.CreateWordParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Word:      word.Word,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := databaseWordToWord(db)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func (apiCfg *apiConfig) handleGetWords(c *fiber.Ctx) error {
	db, err := apiCfg.DB.GetWords(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := databaseWordsToWords(db)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func (apiCfg *apiConfig) handleGetWordById(c *fiber.Ctx) error {
	wordId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := apiCfg.DB.GetWordById(c.Context(), wordId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := databaseWordToWord(db)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}
