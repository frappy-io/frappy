package handlers

import "github.com/gofiber/fiber/v2"

func FrappyHandlers() {

	router := fiber.New()

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is Healthy")
	})

	router.Listen(":8000")
}
