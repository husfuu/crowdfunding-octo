package api

import (
	"github.com/husfuu/crowdfunding-octo/config"
	campaign "github.com/husfuu/crowdfunding-octo/internal/api/campaign/repository"
	transaction "github.com/husfuu/crowdfunding-octo/internal/api/transaction/repository"
	user "github.com/husfuu/crowdfunding-octo/internal/api/user/repository"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type ApiRepository struct {
	Campaign    campaign.Repository
	User        user.Repository
	Transaction transaction.Repository
	// other's domain
}

func NewApiRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) ApiRepository {
	return ApiRepository{
		Campaign:    campaign.NewCampaignRepo(dbList),
		User:        user.NewUserRepo(dbList),
		Transaction: transaction.NewTransactionRepo(dbList),
	}
}
