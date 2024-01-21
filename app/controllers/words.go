package controllers

import (
	"time"

	"github.com/AlvinMrema/kighala-api/app/models"
	dbConfig "github.com/AlvinMrema/kighala-api/platform/database"
	database "github.com/AlvinMrema/kighala-api/platform/database/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateWord func for creates a new word.
// @Description Create a new word.
// @Summary create a new word
// @Tags Words
// @Accept json
// @Produce json
// @Param word body models.Word true "Word"
// @Success 201 {object} models.Word
// @Security ApiKeyAuth
// @Router /kamusi/words [post]
func CreateWord(c *fiber.Ctx) error {
	// Create database connection.
	db, err := dbConfig.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data := models.Word{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Creating the 'word' to database
	dbResult, err := db.DB.CreateWord(c.Context(), database.CreateWordParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Word:      data.Word,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := models.DatabaseWordToWord(dbResult)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

// GetWords func gets all words.
// @Description Get all words.
// @Summary get all words
// @Tags Words
// @Accept json
// @Produce json
// @Success 200 {array} models.Word
// @Router /kamusi/words [get]
func GetWords(c *fiber.Ctx) error {
	// Create database connection.
	db, err := dbConfig.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	dbResult, err := db.DB.GetWords(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := models.DatabaseWordsToWords(dbResult)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

// GetWordById func gets a word associated with ID.
// @Description Get word by ID.
// @Summary get word by ID
// @Tags Words
// @Accept json
// @Produce json
// @Param id path string true "Word ID"
// @Success 200 {object} models.Word
// @Security ApiKeyAuth
// @Router /kamusi/words/{id} [get]
func GetWordById(c *fiber.Ctx) error {
	// Create database connection.
	db, err := dbConfig.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	dbResult, err := db.DB.GetWordById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := models.DatabaseWordToWord(dbResult)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

// UpdateWord func for word updates by given ID.
// @Description Update word.
// @Summary update word
// @Tags Words
// @Accept json
// @Produce json
// @Param id body string true "Word ID"
// @Param word body string true "Word"
// @Success 200 {string} status "ok"
// @Security ApiKeyAuth
// @Router /kamusi/words/{id} [put]
func UpdateWord(c *fiber.Ctx) error {
	// Create database connection.
	db, err := dbConfig.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data := models.Word{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	dbResult, err := db.DB.UpdateWord(c.Context(), database.UpdateWordParams{
		ID:        id,
		UpdatedAt: time.Now().UTC(),
		Word:      data.Word,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := models.DatabaseWordToWord(dbResult)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

// DeleteWord func for deletes word by given ID.
// @Description Delete word by given ID.
// @Summary delete word by given ID
// @Tags Words
// @Accept json
// @Produce json
// @Param id body string true "Word ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /kamusi/words/{id} [delete]
func DeleteWord(c *fiber.Ctx) error {
	// Create database connection.
	db, err := dbConfig.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = db.DB.DeleteWord(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"error":   false,
		"results": "Deleted Successfully",
	})
}
