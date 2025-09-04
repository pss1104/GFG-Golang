package main

import (
	"cars/config"
	"cars/handlers"
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger" ///this is fibre specific middleware, unlike default gorm/ net-http that get added to the handler chain
)

func main() {

	config.ConnectDB()

	app := fiber.New()
	app.Use(logger.New())

	// API endpoints:
	// POST:   http://localhost:5050/cars
	// GET:    http://localhost:5050/cars/:id
	// DELETE: http://localhost:5050/cars/:id

	app.Post("/cars",handlers.CreateCar)
	app.Get("/cars/:id",handlers.GetCar)
	app.Delete("/cars/:id",handlers.DeleteCar)

	fmt.Println("Fiber HTTP Server listening on 5050...")
	if err:=app.Listen(":5050"); err!=nil{
		log.Fatal("Couldnt listen on port 5050",err)
	}
}
