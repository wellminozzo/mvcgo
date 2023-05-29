package controllers

import "github.com/gofiber/fiber/v2"

func HelloIndex(c *fiber.Ctx) error {
	return c.SendString("Hellow world!")
}

func HelloJson(c *fiber.Ctx) error {
	data := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(data)
}
