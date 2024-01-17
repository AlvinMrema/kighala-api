package main

import (
	"log"
	"os"
	// "os"

	"github.com/AlvinMrema/kighala-api/pkg/routes"
	"github.com/AlvinMrema/kighala-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

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

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	// log.Fatal(app.Listen(":3000"))
	utils.StartServer(portString, app)
}
