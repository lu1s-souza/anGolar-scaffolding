package main

import (
	"github.com/gofiber/fiber/v2"
	routes "github.com/lu1s-souza/anGolar-scaffolding/api"
)

func main() {
	app := fiber.New()
	routes.PublicRoutes(app)

	app.Listen(":3000")
}
