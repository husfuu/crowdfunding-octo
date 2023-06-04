package repository

import (
	"github.com/husfuu/crowdfunding-octo/config"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/repository/api"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	// Api api.ApiRepository
	Api api.ApiRepository
}

func NewRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) Repository {
	return Repository{
		Api: api.NewApiRepository(conf, dbList, log),
	}
}
