package main

import "github.com/gofiber/fiber/v2"

type MongoInstance struct {
	Client
	Db
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = ""

type Employee struct {
	ID     string
	Name   string
	Salary float64
	Age    float64
}

func Connect() error {

}

func main() {
	app := fiber.New()

	app.Get("/employee")
	app.Post("/employee")
	app.Put("/employee/:id")
	app.Delete("/employee/:id")
}
