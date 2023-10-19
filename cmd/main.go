package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/greybluesea/order-placement-api-gofiber-gorm-postgres/database"
	"github.com/greybluesea/order-placement-api-gofiber-gorm-postgres/routes"
)

func main() {
	database.ConnectDB()

	app := fiber.New(fiber.Config{})
	app.Get("/", homeHandler)
	routes.SetupCustomerRoutes(app)

	app.Static("/", "./public")
	log.Fatal(app.Listen(":3000"))
}

func homeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, welcome to the Order Placement API 👋!")

}
