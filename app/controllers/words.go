package controllers

import (
	"time"

	"github.com/AlvinMrema/kighala-api/app/models"
	dbConfig "github.com/AlvinMrema/kighala-api/platform/database"
	database "github.com/AlvinMrema/kighala-api/platform/database/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

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

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

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

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"error":   false,
		"results": "Deleted Successfully",
	})
}
