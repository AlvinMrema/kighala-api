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

func (apiCfg *apiConfig) handleCreateWord(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"type":     "https://example.com/probs/out-of-credit",
		"title":    "You do not have enough credit.",
		"status":   403,
		"detail":   "Your current balance is 30, but that costs 50.",
		"instance": "/account/12345/msgs/abc",
	}, "application/problem+json")
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
		AllowOrigins:   "https://*, http://*",
		AllowMethods:   "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:   "*",
		ExposeHeaders:   "Link",
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, Alvin!")
	// })

	// app.Get("/:value", func(c *fiber.Ctx) error {
	// 	return c.SendString("value: " + c.Params("value"))
	// 	// => Get request with value: hello world
	// })

	// app.Get("/:name?", func(c *fiber.Ctx) error {
	// 	if c.Params("name") != "" {
	// 		return c.SendString("Hello " + c.Params("name"))
	// 		// => Hello john
	// 	}
	// 	return c.SendString("Where is john?")
	// })

	app.Get("/test", apiCfg.handleCreateWord)

	app.Listen(":3000")
}
