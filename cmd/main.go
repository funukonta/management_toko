package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := ":8080"
	if os.Getenv("APP_PORT") != "" {
		port = ":" + os.Getenv("APP_PORT")
	}
	app.Listen(port)
}
