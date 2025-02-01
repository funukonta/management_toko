package main

import (
	"os"

	"github.com/funukonta/management_toko/internal/api/middleware"
	"github.com/funukonta/management_toko/internal/api/routes"
	"github.com/funukonta/management_toko/pkg/database"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.ConnectPostgres()
	app := fiber.New()

	app.Use(middleware.ErrorHandler())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.UserRoutes(app, db)
	routes.ProductRoutes(app, db)

	port := ":8080"
	if os.Getenv("APP_PORT") != "" {
		port = ":" + os.Getenv("APP_PORT")
	}
	app.Listen(port)
}
