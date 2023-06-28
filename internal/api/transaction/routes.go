package transaction

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/handler"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Get("/transactions", handler.Api.Transaction.GetTransactionById)
	api.Get("/transactions/user", handler.Api.Transaction.GetTransactionByUserId)
	api.Get("/transactions/campaign", handler.Api.Transaction.GetTransactionByCampaignId)
}
