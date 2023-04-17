package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husfuu/crowdfunding-octo/app/handlers"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Get("/campaigns", handlers.GetCampaigns)
	route.Get("/campaigns/:id", handlers.GetCampaign)
	route.Post("/campaigns", handlers.CreateCampaign)
}
