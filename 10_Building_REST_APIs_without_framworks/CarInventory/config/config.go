package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// connect to postgres database
func ConnectDB() {
	dsn := "host=localhost port=5432 user=postgres password=Saibaba@1 dbname=mydb sslmode=disable"
	db, error := sql.Open("postgres", dsn)
	if error != nil {
		fmt.Println("Error while connecting to database")
		panic(error)
	}
	if error = db.Ping(); error != nil {
		fmt.Println("Error while pinging to database")
		panic(error)
	}
	fmt.Println("Connected to database successfully")
	DB = db
}
