package main

import (
	"cars/config"
	"cars/handlers"
	"fmt"
	"net/http"
)

func main() {
	
	config.ConnectDB()
	
	http.HandleFunc("/cars", handlers.CarHandler)
	http.HandleFunc("/cars/", handlers.CarHandler)

	fmt.Println("Server listening on 5050...")
	http.ListenAndServe(":5050", nil)
}