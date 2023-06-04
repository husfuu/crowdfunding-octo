package handler

import (
	"github.com/husfuu/crowdfunding-octo/config"
	api "github.com/husfuu/crowdfunding-octo/internal/wrapper/handler/api"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/usecase"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Api api.ApiHandler
}

func NewHandler(usecase usecase.Usecase, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) Handler {
	return Handler{
		Api: api.NewApiHandler(usecase, conf, log),
	}
}
