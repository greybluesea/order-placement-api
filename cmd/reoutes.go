package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/greybluesea/order-placement-api-gofiber-gorm-postgres/database"
	"github.com/greybluesea/order-placement-api-gofiber-gorm-postgres/models"
	//"github.com/google/uuid"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", homeHandler)
	//app.Get("/customers", getAllCustomers)
	app.Post("/customers/new", createCustomer)
	//app.Post("/customers/delete/:id", deleteCustomerByID)
	//app.Get("/customers/edit/:id", editCustomerByID)
	//app.Post("/edit/:id", updateCustomerByID)
}

func homeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, welcome to the Order Placement API!")

}

func newfactHandler(c *fiber.Ctx) error {
	return nil
}

func createCustomer(c *fiber.Ctx) error {
	var customer models.Customer

	if err := c.BodyParser(&customer); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&customer)
	return c.Status(200).JSON(customer)
}

func editHandler(c *fiber.Ctx) error {
	return nil
}

func updateHandler(c *fiber.Ctx) error {
	return nil
}

func deleteHandler(c *fiber.Ctx) error {
	return nil
}

/* func helloHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World! Tony here 👋!!")
} */
