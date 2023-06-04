package api

import (
	"github.com/husfuu/crowdfunding-octo/config"
	campaign "github.com/husfuu/crowdfunding-octo/internal/api/campaign/usecase"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/repository"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type ApiUsecase struct {
	Campaign campaign.CampaignUsecase
}

func NewApiUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) ApiUsecase {
	return ApiUsecase{
		Campaign: campaign.NewCampaignUsecase(repo, conf, dbList, log),
	}
}
