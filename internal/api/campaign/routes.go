package campaign

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/handler"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	// api.Get("/campaigns", handler.Api.Campaign.GetListCampaigns)
	api.Get("/campaign/:id", handler.Api.Campaign.GetCampaign)
	api.Post("/campaign", handler.Api.Campaign.PostCampaign)
	api.Post("/campaign-image", handler.Api.Campaign.PostCImage)
	api.Put("/campaign/:id", handler.Api.Campaign.UpdateCampaign)
}
