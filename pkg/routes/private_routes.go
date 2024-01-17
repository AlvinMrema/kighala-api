package routes

import (
	"github.com/AlvinMrema/kighala-api/app/controllers"
	"github.com/AlvinMrema/kighala-api/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	dictionary := route.Group("/kamusi")

	dictionary.Post("/words", middlewares.AuthMiddleware(), controllers.CreateWord)
	dictionary.Get("/words/:id", middlewares.AuthMiddleware(), controllers.GetWordById)
	dictionary.Put("/words/:id", middlewares.AuthMiddleware(), controllers.UpdateWord)
	dictionary.Delete("/words/:id", middlewares.AuthMiddleware(), controllers.DeleteWord)

	dictionary.Get("/definitions", middlewares.AuthMiddleware(), controllers.GetDefinitions)
	dictionary.Post("/definitions", middlewares.AuthMiddleware(), controllers.CreateDefinition)
	dictionary.Get("/definitions/:id", middlewares.AuthMiddleware(), controllers.GetDefinitionById)
	dictionary.Put("/definitions/:id", middlewares.AuthMiddleware(), controllers.UpdateDefinition)
	dictionary.Delete("/definitions/:id", middlewares.AuthMiddleware(), controllers.DeleteDefinition)
}
