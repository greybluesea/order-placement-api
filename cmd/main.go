package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/greybluesea/order-processing-api-gofiber-gorm-postgres/database"
)

func main() {
	database.ConnectDB()

	app := fiber.New(fiber.Config{})

	app.Get("/", homeHandler)
	app.Get("/newfact", newfactHandler)
	app.Post("/create", createHandler)
	app.Post("/delete", deleteHandler)
	app.Get("/edit/:id", editHandler)
	app.Post("/edit/:id", updateHandler)

	app.Static("/", "./public")
	app.Listen(":3000")
}
