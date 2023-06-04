package usecase

import (
	"github.com/husfuu/crowdfunding-octo/config"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/repository"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/usecase/api"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	Api api.ApiUsecase
}

func NewUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) Usecase {
	return Usecase{
		Api: api.NewApiUsecase(repo, conf, dbList, log),
	}
}
