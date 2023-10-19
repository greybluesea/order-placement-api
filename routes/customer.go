package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/greybluesea/order-placement-api-gofiber-gorm-postgres/database"
	"github.com/greybluesea/order-placement-api-gofiber-gorm-postgres/models"
	//"github.com/google/uuid"
)

func SetupCustomersRoutes(app *fiber.App) {

	app.Get("/customers", getAllCustomers)
	app.Post("/customers/new", createCustomer)
	app.Get("/customers/:id", getCustomerByID)
	app.Post("/customers/:id", updateCustomerByID)
	app.Delete("/customers/:id", deleteCustomerByID)
}

func getAllCustomers(c *fiber.Ctx) error {
	customers := []models.Customer{}
	database.DB.Find(&customers)

	if len(customers) == 0 {
		return c.Status(200).JSON("No customer in database yet.")
	}

	return c.Status(200).JSON(customers)
}

func createCustomer(c *fiber.Ctx) error {
	customer := models.Customer{}

	if err := c.BodyParser(&customer); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&customer)
	return c.Status(200).JSON(customer)
}

func getCustomerByID(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	customer := models.Customer{}
	database.DB.Find(&customer, "id = ?", id)
	if customer.ID == 0 {
		return c.Status(400).JSON("user does not exist")
	}

	return c.Status(200).JSON(customer)
}

func updateCustomerByID(c *fiber.Ctx) error {

	//find customer by id from url params
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	customer := models.Customer{}
	database.DB.Find(&customer, "id = ?", id)
	if customer.ID == 0 {
		return c.Status(400).JSON("user does not exist")
	}

	//update customer by request body
	customerName := models.CustomerName{}
	if err := c.BodyParser(&customerName); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	customer.FirstName = customerName.FirstName
	customer.LastName = customerName.LastName
	database.DB.Save(&customer)

	return c.Status(200).JSON(customer)
}

type DeleteCustomerResponse struct {
	Message  string          `json:"message"`
	Customer models.Customer `json:"customer"`
}

func deleteCustomerByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	customer := models.Customer{}
	database.DB.Find(&customer, "id = ?", id)
	if customer.ID == 0 {
		return c.Status(400).JSON("user does not exist")
	}

	if err = database.DB.Delete(&customer).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	response := DeleteCustomerResponse{
		Message:  "Successfully deleted customer: ",
		Customer: customer,
	}

	//	return c.Status(200).JSON(fmt.Sprintf("Successfully deleted customer id: %d", id))
	return c.Status(200).JSON(response)
}
