package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	// txRoutes "github.com/siddhantprateek/frappy/handlers/merger"
	controllers "github.com/siddhantprateek/frappy/handlers/controllers"
)

func FrappyHandlers() {

	router := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Cross-Origin Resource Sharing: CORS middleware
	router.Use(cors.New())

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is Healthy")
	})

	router.Get("/block", controllers.GetBlock)
	router.Get("/account", controllers.GetAccount)
	router.Get("/events", controllers.GetEvents)
	router.Get("/transaction", controllers.GetTransaction)
	router.Get("/build", controllers.BuildTransaction)

	// router.Use("/api/transaction", router)

	router.Listen(":8000")
}
