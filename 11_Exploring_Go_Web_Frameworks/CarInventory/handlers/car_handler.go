package handlers

import (
	"cars/models"
	"fmt"
	"strconv"
	"sync"
	"github.com/gofiber/fiber/v2"
)
var mu sync.Mutex
// create car API
func CreateCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}

	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON, incorrect input body",
			"detail": err.Error(),
		})
	}

	car.Insert()
	
	fmt.Println("Car saved to the inventory with id:", car.ID)

	return c.Status(fiber.StatusCreated).JSON(car)
}

// get car API
// for future reference: using _ instead of req to maintain functino signature
func GetCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}

	id,err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid car ID",
			"detail": err.Error(),
		})
	}
	car.ID = id
	if err := car.Get(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Car with given id not found",
			"id": car.ID,
		})
	}
	fmt.Println("Fetched car:", car)

	return c.Status(fiber.StatusOK).JSON(car)
}

// delete car API
func DeleteCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	id,err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid car ID",
			"detail": err.Error(),
		})
	}
	//getting the car with required id
	car := &models.Car{}
	car.ID = id

	car.Delete()
	
	fmt.Println("Deleted car:", car)
	
	return c.SendStatus(fiber.StatusNoContent)
}

//sendstatus vs status
// sendstatus sends only the status code
// status sends the status code and allows sending a body with it