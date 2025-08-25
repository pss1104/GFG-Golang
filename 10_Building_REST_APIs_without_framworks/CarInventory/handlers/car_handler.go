package handlers

import (
	"cars/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var mu sync.Mutex

// API endpoints:
// GET:    http://localhost:5050/cars/:id
// POST:   http://localhost:5050/cars
// DELETE: http://localhost:5050/cars/:id

func CarHandler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	id := strings.TrimPrefix(path, "/cars")
	id = strings.Trim(id, "/")

	switch req.Method {
	case "POST":
		if id == "" {
			createCar(res, req)
		} else {
			http.Error(res, "Invalid POST request", http.StatusBadRequest)
		}
	case "GET":
		if id == "" {
			http.Error(res, "Car ID is required", http.StatusBadRequest)
		} else {
			carID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(res, "Invalid car ID", http.StatusBadRequest)
				return
			}
			getCar(res, req, carID)
		}
	case "DELETE":
		if id == "" {
			http.Error(res, "Car ID is required", http.StatusBadRequest)
		} else {
			carID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(res, "Invalid car ID", http.StatusBadRequest)
				return
			}
			deleteCar(res, req, carID)
		}
	default:
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// create car API
func createCar(res http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}

	if err := json.NewDecoder(req.Body).Decode(&car); err != nil {
		http.Error(res, "Incorrect JSON input", http.StatusBadRequest)
		return
	}

	car.Insert()
	models.Cars[int(car.ID)] = *car
	fmt.Println("Created car:", car)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(car)
}

// get car API
// for future reference: using _ instead of req to maintain functino signature
func getCar(res http.ResponseWriter, _ *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	//getting the car with required id
	car := &models.Car{ID: id}
	//alternative: use a parameter id in the Get function car.Get(id)

	if err:=car.Get(); err != nil {
		fmt.Println("Error while getting car from database")
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		return
	}


	fmt.Println("Retrieved car:", car)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(car)
}

// delete car API
func deleteCar(res http.ResponseWriter, _ *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	//getting the car with required id
	car := &models.Car{}
	car.ID = id

	car.Delete()
	
	fmt.Println("Deleted car:", car)
	
	res.WriteHeader(http.StatusNoContent)
}
