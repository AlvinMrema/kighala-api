package main

import (
	"time"

	"github.com/AlvinMrema/kighala-api/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateWord(c *fiber.Ctx) error {
	data := Word{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Creating the 'word' to database
	db, err := apiCfg.DB.CreateWord(c.Context(), database.CreateWordParams{
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func (apiCfg *apiConfig) handleGetWordById(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := apiCfg.DB.GetWordById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := databaseWordToWord(db)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func (apiCfg *apiConfig) handleUpdateWord(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data := Word{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := apiCfg.DB.UpdateWord(c.Context(), database.UpdateWordParams{
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

	results := databaseWordToWord(db)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func (apiCfg *apiConfig) handleDeleteWord(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = apiCfg.DB.DeleteWord(c.Context(), id)
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

func (apiCfg *apiConfig) handleGetDefinitionsByWordId(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := apiCfg.DB.GetDefinitionsByWordID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := databaseDefinitionsToDefinitions(db)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func (apiCfg *apiConfig) handleCreateDefinition(c *fiber.Ctx) error {
	data := Definition{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Creating the 'definition' to database
	db, err := apiCfg.DB.CreateDefinition(c.Context(), database.CreateDefinitionParams{
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

	results := databaseDefinitionToDefinition(db)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func (apiCfg *apiConfig) handleGetDefinitions(c *fiber.Ctx) error {
	db, err := apiCfg.DB.GetDefinitions(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := databaseDefinitionsToDefinitions(db)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func (apiCfg *apiConfig) handleGetDefinitionById(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := apiCfg.DB.GetDefinitionById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	results := databaseDefinitionToDefinition(db)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func (apiCfg *apiConfig) handleUpdateDefinition(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data := Definition{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := apiCfg.DB.UpdateDefinition(c.Context(), database.UpdateDefinitionParams{
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

	results := databaseDefinitionToDefinition(db)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"results": results,
	})
}

func (apiCfg *apiConfig) handleDeleteDefinition(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = apiCfg.DB.DeleteDefinition(c.Context(), id)
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