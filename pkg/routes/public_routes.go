package routes

import (
	"github.com/AlvinMrema/kighala-api/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	// route := a.Group("/api/v1")
	route := a.Group("/api/v1")

	auth := route.Group("/auth")
	auth.Post("/signin", controllers.SignIn)
	auth.Post("signup", controllers.SignUp)

	dictionary := route.Group("/kamusi")
	dictionary.Get("/words", controllers.GetWords)
	dictionary.Get("/words/:id/definitions", controllers.GetDefinitionsByWordId)
}
