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

	_ "github.com/AlvinMrema/kighala-api/docs" // load API Docs files (Swagger)
)

// @title Kighala API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email sonalpha023@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 0.0.0.0:3000
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Auth-Token
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

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://kighala-api-production.up.railway.app/, http://*",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept",
		// ExposeHeaders:    "Link",
		AllowCredentials: false,
		MaxAge:           300,
	}))
	// app.Use(cors.New())

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// log.Fatal(app.Listen(":3000"))
	utils.StartServer(portString, app)
}
