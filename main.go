package main

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type Car struct {
	Name string
	Price float64
}

// slice: like an infinit array
var cars []Car

func generateCars () {
	cars = append(cars, Car{Name: "Ferrari", Price: 100})
	cars = append(cars, Car{Name: "Mercedes", Price: 200})
	cars = append(cars, Car{Name: "Porsche", Price: 300})
}

func main() {
	generateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":8080"))
}


func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createCar (c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	saveCar(*car)
	return c.JSON(200, cars)
}

func saveCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars (name, price) VALUES ($1, $2)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(car.Name, car.Price)
	if err != nil {
		return err
	}
	return nil
}