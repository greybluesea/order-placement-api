package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/greybluesea/order-placement-api-gofiber-gorm-postgres/database"
)

func main() {
	database.ConnectDB()

	app := fiber.New(fiber.Config{})
	setupRoutes(app)

	app.Static("/", "./public")
	log.Fatal(app.Listen(":3000"))
}
