package api

import (
	"github.com/husfuu/crowdfunding-octo/config"
	campaign "github.com/husfuu/crowdfunding-octo/internal/api/campaign/repository"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type ApiRepository struct {
	Campaign campaign.Repository
	// other's domain
}

func NewApiRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) ApiRepository {
	return ApiRepository{
		Campaign: campaign.NewCampaignRepo(dbList),
	}
}
