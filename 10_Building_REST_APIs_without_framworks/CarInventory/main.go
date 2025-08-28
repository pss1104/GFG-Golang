package main

import (
	"cars/config"
	"cars/handlers"
	middlewares "cars/middleware"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	config.ConnectDB()
	mux := mux.NewRouter()

	mux.HandleFunc("/cars", handlers.CarHandler)
	mux.HandleFunc("/cars/{id}", handlers.CarHandler)

	wrappedmux := middlewares.Logger(mux)
	wrappedmux = middlewares.SecurityHeaders(wrappedmux)

	fmt.Println("Server listening on 5050...")
	http.ListenAndServe(":5050", wrappedmux)
}
