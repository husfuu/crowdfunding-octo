package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husfuu/crowdfunding-octo/pkg/routes"
)

func main() {
	app := fiber.New()

	routes.PublicRoutes(app)

	app.Listen(":3000")
}
