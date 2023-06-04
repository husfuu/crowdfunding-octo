package api

import (
	"github.com/husfuu/crowdfunding-octo/config"
	campaign "github.com/husfuu/crowdfunding-octo/internal/api/campaign/delivery"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/usecase"
	"github.com/sirupsen/logrus"
)

type ApiHandler struct {
	Campaign campaign.CampaignHandler
}

func NewApiHandler(usecase usecase.Usecase, conf *config.Config, log *logrus.Logger) ApiHandler {
	return ApiHandler{
		Campaign: campaign.NewCampaignHandler(usecase, conf, log),
	}
}
