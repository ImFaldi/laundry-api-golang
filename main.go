package main

import (
	"github.com/gofiber/fiber"
	"github.com/indrabpn12/laundry-api-golang.git/database"
	"github.com/indrabpn12/laundry-api-golang.git/route"
)

func main() {

	database.Conn()

	app := fiber.New()

	route.Init(app)

	app.Listen(":3000")
}
