package main

import (
	"github.com/gofiber/fiber/v2"
	"gofiber-mysql/database"
	"gofiber-mysql/services"
	"log"
)

func main() {
	if err := database.GetConnection(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Get("/students", services.GetStudents)
	app.Post("/student", services.PostStudent)
	app.Put("/student", services.PutStudent)
	app.Delete("/student", services.DeleteStudent)
	log.Fatal(app.Listen(":8081"))
}
