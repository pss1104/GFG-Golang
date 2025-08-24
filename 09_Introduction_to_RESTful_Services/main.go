package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Car struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// Car Inventory
var Cars = make(map[int]Car)

var mu sync.Mutex

// API endpoints:
// GET:    http://localhost:5050/cars/:id
// POST:   http://localhost:5050/cars
// DELETE: http://localhost:5050/cars/:id

func carHandler(res http.ResponseWriter, req *http.Request) {
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

	var car Car

	if err := json.NewDecoder(req.Body).Decode(&car); err != nil {
		http.Error(res, "Incorrect JSON input", http.StatusBadRequest)
		return
	}

	id := rand.Intn(100)
	car.ID = int32(id)
	Cars[id] = car

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

	car, ok := Cars[id]
	if !ok {
		http.Error(res, "No such car available", http.StatusNotFound)
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

	_, ok := Cars[id]
	if !ok {
		http.Error(res, "No such car available", http.StatusNotFound)
		return
	}

	fmt.Println("Deleted car:", Cars[id])

	delete(Cars, id)
	res.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/cars", carHandler)
	http.HandleFunc("/cars/", carHandler)

	fmt.Println("Server listening on 5050...")
	http.ListenAndServe(":5050", nil)
}
