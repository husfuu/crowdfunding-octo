package api

import (
	"github.com/husfuu/crowdfunding-octo/config"
	campaign "github.com/husfuu/crowdfunding-octo/internal/api/campaign/delivery"
	transaction "github.com/husfuu/crowdfunding-octo/internal/api/transaction/delivery"
	user "github.com/husfuu/crowdfunding-octo/internal/api/user/delivery"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/usecase"
	"github.com/sirupsen/logrus"
)

type ApiHandler struct {
	Campaign    campaign.CampaignHandler
	User        user.UserHandler
	Transaction transaction.TransactionHandler
}

func NewApiHandler(usecase usecase.Usecase, conf *config.Config, log *logrus.Logger) ApiHandler {
	return ApiHandler{
		Campaign:    campaign.NewCampaignHandler(usecase, conf, log),
		User:        user.NewUserHandler(usecase, conf, log),
		Transaction: transaction.NewTransactionHandler(usecase, conf, log),
	}
}
