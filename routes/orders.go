package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/greybluesea/order-placement-api/database"
	"github.com/greybluesea/order-placement-api/models"
)

func SetupOrdersRoutes(app *fiber.App) {
	app.Get("/orders", getAllOrders)
	app.Post("/orders/new", createOrder)
	app.Get("/orders/:id", getOrderByID)
	app.Post("/orders/:id", updateOrderByID)
	app.Delete("/orders/:id", deleteOrderByID)
}

// Handlers for "/orders" routes

func getAllOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.DB.Preload("Customer").Preload("Product").Find(&orders)

	if len(orders) == 0 {
		return c.Status(200).JSON("No orders in the database yet.")
	}

	return c.Status(200).JSON(orders)
}

func createOrder(c *fiber.Ctx) error {
	// Parse the request body to extract CustomerID and ProductID.
	order := models.Order{}

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON("Invalid request body")
	}

	// Retrieve the customer and product using database.DB.Find() directly.
	if err := database.DB.Find(&order.Customer, "id = ?", order.CustomerID).Error; err != nil {
		return c.Status(400).JSON("Failed to retrieve customer")
	}

	if err := database.DB.Find(&order.Product, "id = ?", order.ProductID).Error; err != nil {
		return c.Status(400).JSON("Failed to retrieve product")
	}

	// Save the new order to the database.
	if err := database.DB.Create(&order).Error; err != nil {
		return c.Status(400).JSON(fmt.Sprintln("Failed to create the order", err.Error()))
	}

	return c.Status(200).JSON(order)
}

func getOrderByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	order := models.Order{}
	if err := database.DB.Preload("Customer").Preload("Product").Find(&order, "id = ?", id).Error; err != nil {
		return c.Status(400).JSON("Order does not exist")
	}

	return c.Status(200).JSON(order)
}

func updateOrderByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	order := models.Order{}
	if err := database.DB.Find(&order, "id = ?", id).Error; err != nil {
		return c.Status(400).JSON("Order does not exist")
	}

	// Parse the request body to update the order.

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON("Invalid request body")
	}

	// Retrieve the customer and product using database.DB.Find() directly.
	if err := database.DB.Find(&order.Customer, "id = ?", order.CustomerID).Error; err != nil {
		return c.Status(400).JSON("Failed to retrieve customer")
	}

	if err := database.DB.Find(&order.Product, "id = ?", order.ProductID).Error; err != nil {
		return c.Status(400).JSON("Failed to retrieve product")
	}

	// Save the new order to the database.
	if err := database.DB.Save(&order).Error; err != nil {
		return c.Status(400).JSON("Failed to update the order")
	}

	return c.Status(200).JSON(order)
}

type DeleteOrderResponse struct {
	Message string       `json:"message"`
	Order   models.Order `json:"order"`
}

func deleteOrderByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	order := models.Order{}
	if err := database.DB.Find(&order, "id = ?", id).Error; err != nil {
		return c.Status(400).JSON("Order does not exist")
	}

	if err := database.DB.Delete(&order).Error; err != nil {
		return c.Status(400).JSON("Failed to delete the order")
	}

	response := DeleteOrderResponse{
		Message: "Successfully deleted order",
		Order:   order,
	}

	return c.Status(200).JSON(response)
}
