package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/AlvinMrema/kighala-api/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

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

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://*, http://*",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "*",
		ExposeHeaders:    "Link",
		AllowCredentials: false,
		MaxAge:           300,
	}))

	api := app.Group("/api")

	dictionary := api.Group("/kamusi")
	dictionary.Get("/words", apiCfg.handleGetWords)
	dictionary.Post("/words", apiCfg.handleCreateWord)
	dictionary.Get("/words/:id", apiCfg.handleGetWordById)
	dictionary.Put("/words/:id", apiCfg.handleUpdateWord)
	dictionary.Delete("/words/:id", apiCfg.handleDeleteWord)
	dictionary.Get("/words/:id/definitions", apiCfg.handleGetDefinitionsByWordId)
	
	dictionary.Get("/definitions", apiCfg.handleGetDefinitions)
	dictionary.Post("/definitions", apiCfg.handleCreateDefinition)
	dictionary.Get("/definitions/:id", apiCfg.handleGetDefinitionById)
	dictionary.Put("/definitions/:id", apiCfg.handleUpdateDefinition)
	dictionary.Delete("/definitions/:id", apiCfg.handleDeleteDefinition)

	log.Fatal(app.Listen(":3000"))
}
