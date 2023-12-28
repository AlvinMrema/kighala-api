package main

import (
	// "database/sql"
	"log"
	"os"

	// "github.com/AlvinMrema/kighala-api/internal/controllers"
	// "github.com/AlvinMrema/kighala-api/internal/database"
	"github.com/AlvinMrema/kighala-api/app/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)



func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	// conn, err := sql.Open("postgres", dbURL)
	// if err != nil {
	// 	log.Fatal("Can't connect to database:", err)
	// }

	// apiCfg := apiConfig{
	// 	DB: database.New(conn),
	// }

	app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "https://*, http://*",
	// 	AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
	// 	AllowHeaders:     "*",
	// 	ExposeHeaders:    "Link",
	// 	AllowCredentials: false,
	// 	MaxAge:           300,
	// }))
	app.Use(cors.New())

	api := app.Group("/api")

	dictionary := api.Group("/kamusi")
	dictionary.Get("/words", controllers.GetWords)
	dictionary.Post("/words", controllers.CreateWord)
	dictionary.Get("/words/:id", controllers.GetWordById)
	dictionary.Put("/words/:id", controllers.UpdateWord)
	dictionary.Delete("/words/:id", controllers.DeleteWord)
	dictionary.Get("/words/:id/definitions", controllers.GetDefinitionsByWordId)

	dictionary.Get("/definitions", controllers.GetDefinitions)
	dictionary.Post("/definitions", controllers.CreateDefinition)
	dictionary.Get("/definitions/:id", controllers.GetDefinitionById)
	dictionary.Put("/definitions/:id", controllers.UpdateDefinition)
	dictionary.Delete("/definitions/:id", controllers.DeleteDefinition)

	log.Fatal(app.Listen(":3000"))
}
