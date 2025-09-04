package models

import (
	"cars/config"
	"database/sql"
	"fmt"
)

type Car struct {
	ID   int  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
}

// Car Inventory
var Cars = make(map[int]Car)

// insert function to add cars to DB can be added here
func (c *Car) Insert() {
	query := `INSERT INTO cars (id, name) VALUES ($1, $2) RETURNING id`
	if err := config.DB.QueryRow(query, c.ID, c.Name).Scan(&c.ID); err != nil {
		fmt.Println("Error while inserting car into database")
		panic(err)
	}
}

// update function
func (c *Car) Update() {
	query := `UPDATE cars SET name=$1 WHERE id=$2`
	if _, err := config.DB.Exec(query, c.Name, c.ID); err != nil {
		fmt.Println("Error while updating car in database")
		panic(err)
	}
}

// get function
func (c *Car) Get() error {
	query := `SELECT id, name FROM cars WHERE id=$1`
	if err := config.DB.QueryRow(query, c.ID).Scan(&c.ID, &c.Name); err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Car with ID %d not found", c.ID)
			return err
		}
	}
	return nil
}

// delete function
func (c *Car) Delete() {
	query := `DELETE FROM cars WHERE id=$1`
	if _, err := config.DB.Exec(query, c.ID); err != nil {
		fmt.Println("Error while deleting car from database")
		panic(err)
	}
}