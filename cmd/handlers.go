package main

import (
	"github.com/gofiber/fiber/v2"
)

func homeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")

}

func newfactHandler(c *fiber.Ctx) error {
	return nil
}

func createHandler(c *fiber.Ctx) error {
	return nil
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
