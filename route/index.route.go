package route

import (
	"github.com/gofiber/fiber"
	"github.com/indrabpn12/laundry-api-golang.git/controller"
)

func Init(app *fiber.App) {
	app.Get("/api/customer", func(c *fiber.Ctx) {
		controller.GetCustomerAll(c)
	})
}
