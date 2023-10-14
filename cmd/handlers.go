package main

import (
	"github.com/gofiber/fiber/v2"
)

/*
	 func SetupRoutes(app *fiber.App) {
		//	app.Get("/hello", helloHandler)
		app.Get("/", homeHandler)
		app.Get("/newfact", newfactHandler)
		app.Post("/create", createHandler)
		app.Post("/delete", deleteHandler)
		app.Get("/edit/:id", editHandler)
		app.Post("/edit/:id", updateHandler)
	}
*/
func homeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func newfactHandler(c *fiber.Ctx) error {
	return nil
}

func createHandler(c *fiber.Ctx) error {
	return nil
}

/* func successHandler(c *fiber.Ctx) error {
	return c.Render("success", fiber.Map{
		"Title": "Fact added successfully",
	})
} */

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
