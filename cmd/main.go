package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/greybluesea/order-placement-api/database"
	"github.com/greybluesea/order-placement-api/routes"
	/* "github.com/joho/godotenv" */)

func main() {
	/* 	err := godotenv.Load()
	   	if err != nil {
	   		log.Fatal("Error loading .env file")
	   	} */

	database.ConnectDB()

	app := fiber.New(fiber.Config{})
	app.Get("/", homeHandler)
	routes.SetupCustomersRoutes(app)
	routes.SetupProductsRoutes(app)
	routes.SetupOrdersRoutes(app)

	app.Static("/", "./public")
	log.Fatal(app.Listen(":3000"))
}

func homeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, welcome to the Order Placement API 👋!")

}
