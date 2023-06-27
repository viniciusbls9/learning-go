package main

import "github.com/labstack/echo/v4"

type Car struct {
	Name string
	Price float64
}

// slice: like an infinit array
var cars []Car

func createCars () {
	cars = append(cars, Car{Name: "Ferrari", Price: 100})
	cars = append(cars, Car{Name: "Mercedes", Price: 200})
	cars = append(cars, Car{Name: "Porsche", Price: 300})
}

func main() {
	createCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.Logger.Fatal(e.Start(":8080"))
}


func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}