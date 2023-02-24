package merger

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/siddhantprateek/frappy/handlers/controllers"
)

func TransactionAPI() *fiber.App {

	txRouter := fiber.New()

	txRouter.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Transaction route is Healthy ✅")
	})
	txRouter.Get("/block", controllers.GetBlock)
	txRouter.Get("/account", controllers.GetAccount)
	txRouter.Get("/events", controllers.GetEvents)
	txRouter.Get("/transaction", controllers.GetTransaction)
	txRouter.Get("/build", controllers.BuildTransaction)

	return txRouter
}
