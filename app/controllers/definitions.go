package controllers

import (
	"time"

	"github.com/AlvinMrema/kighala-api/app/models"
	dbConfig "github.com/AlvinMrema/kighala-api/platform/database"
	database "github.com/AlvinMrema/kighala-api/platform/database/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetDefinitionsByWordId(c *fiber.Ctx) error {
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

	dbResult, err := db.DB.GetDefinitionsByWordID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := models.DatabaseDefinitionsToDefinitions(dbResult)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func CreateDefinition(c *fiber.Ctx) error {
	// Create database connection.
	db, err := dbConfig.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data := models.Definition{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Creating the 'definition' to database
	dbResult, err := db.DB.CreateDefinition(c.Context(), database.CreateDefinitionParams{
		ID:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		WordID:       data.WordID,
		Definition:   data.Definition,
		PartOfSpeech: data.PartOfSpeech,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := models.DatabaseDefinitionToDefinition(dbResult)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func GetDefinitions(c *fiber.Ctx) error {
	// Create database connection.
	db, err := dbConfig.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	dbResult, err := db.DB.GetDefinitions(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := models.DatabaseDefinitionsToDefinitions(dbResult)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func GetDefinitionById(c *fiber.Ctx) error {
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

	dbResult, err := db.DB.GetDefinitionById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := models.DatabaseDefinitionToDefinition(dbResult)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func UpdateDefinition(c *fiber.Ctx) error {
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

	data := models.Definition{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	dbResult, err := db.DB.UpdateDefinition(c.Context(), database.UpdateDefinitionParams{
		ID:           id,
		UpdatedAt:    time.Now().UTC(),
		Definition:   data.Definition,
		PartOfSpeech: data.PartOfSpeech,
		WordID:       data.WordID,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := models.DatabaseDefinitionToDefinition(dbResult)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func DeleteDefinition(c *fiber.Ctx) error {
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

	err = db.DB.DeleteDefinition(c.Context(), id)
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
