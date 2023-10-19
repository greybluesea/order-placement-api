package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/greybluesea/order-placement-api-gofiber-gorm-postgres/database"
	"github.com/greybluesea/order-placement-api-gofiber-gorm-postgres/models"
)

func SetupProductsRoutes(app *fiber.App) {
	app.Get("/products", getAllProducts)
	app.Post("/products/new", createProduct)
	app.Get("/products/:id", getProductByID)
	app.Post("/products/:id", updateProductByID)
	app.Delete("/products/:id", deleteProductByID)
}

// Handlers for "/products" routes

func getAllProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.DB.Find(&products)

	if len(products) == 0 {
		return c.Status(200).JSON("No products in the database yet.")
	}

	return c.Status(200).JSON(products)
}

func createProduct(c *fiber.Ctx) error {
	product := models.Product{}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&product)
	return c.Status(200).JSON(product)
}

func getProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	product := models.Product{}
	database.DB.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return c.Status(400).JSON("Product does not exist")
	}

	return c.Status(200).JSON(product)
}

func updateProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	product := models.Product{}
	database.DB.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return c.Status(400).JSON("Product does not exist")
	}

	productData := models.Product{}
	if err := c.BodyParser(&productData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	product.Name = productData.Name
	product.Price = productData.Price
	database.DB.Save(&product)

	return c.Status(200).JSON(product)
}

type DeleteProductResponse struct {
	Message string         `json:"message"`
	Product models.Product `json:"product"`
}

func deleteProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	product := models.Product{}
	database.DB.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return c.Status(400).JSON("Product does not exist")
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	response := DeleteProductResponse{
		Message: "Successfully deleted product:",
		Product: product,
	}

	return c.Status(200).JSON(response)
}
