package main

import (
	"cars/config"
	"cars/handlers"
	middlewares "cars/middleware"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger" ///this is fibre specific middleware, unlike default gorm/ net-http that get added to the handler chain
)

func main() {

	config.ConnectDB()

	app := fiber.New()
	app.Use(logger.New())

	//middlewares
	app.Use(middlewares.SecurityHeaders)

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "admin123",
			"john":  "doe",
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	app.Use(etag.New())
	app.Use(cors.New())

	// API endpoints:
	// POST:   http://localhost:5050/cars
	// GET:    http://localhost:5050/cars/:id
	// DELETE: http://localhost:5050/cars/:id

	app.Post("/cars", handlers.CreateCar)
	app.Get("/cars/:id", handlers.GetCar)
	app.Delete("/cars/:id", handlers.DeleteCar)

	fmt.Println("Fiber HTTP Server listening on 5050...")
	if err := app.Listen(":5050"); err != nil {
		log.Fatal("Couldnt listen on port 5050", err)
	}
}
