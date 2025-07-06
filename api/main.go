package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"your_module_name/routes" // Replace with your actual module name
)

func setRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	// Middleware to redirect all HTTP requests to HTTPS
	app.Use(func(c *fiber.Ctx) error {
		if c.Protocol() == "http" {
			return c.Redirect("https://" + c.Hostname() + c.OriginalURL(), fiber.StatusPermanentRedirect)
		}
		return c.Next()
	})

	//  Request logger
	app.Use(logger.New())

	setRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
