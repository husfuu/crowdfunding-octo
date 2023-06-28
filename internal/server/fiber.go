package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/husfuu/crowdfunding-octo/config"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/handler"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/repository"
	"github.com/husfuu/crowdfunding-octo/internal/wrapper/usecase"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/db"
	"github.com/sirupsen/logrus"

	apiCampaign "github.com/husfuu/crowdfunding-octo/internal/api/campaign"
	apiTransaction "github.com/husfuu/crowdfunding-octo/internal/api/transaction"
	apiUser "github.com/husfuu/crowdfunding-octo/internal/api/user"
)

func Run(conf *config.Config, dbList *db.DatabaseList, appLoger *logrus.Logger) {
	app := fiber.New(fiber.Config{
		AppName:      conf.App.Name,
		ServerHeader: "Gooooo",
		// Views:        engine,
		BodyLimit: conf.App.BodyLimit * 1024 * 1024,
	})

	// wrapper results
	repo := repository.NewRepository(conf, dbList, appLoger)
	usecase := usecase.NewUsecase(repo, conf, dbList, appLoger)
	// usecase := usecase.NewCampaignUsecase(repo, conf, dbList, appLoger)
	handler := handler.NewHandler(usecase, conf, dbList, appLoger)

	//* api endpoint
	api := app.Group(conf.App.Endpoint)

	// Routes
	apiCampaign.NewRoutes(api, handler)
	apiUser.NewRoutes(api, handler)
	apiTransaction.NewRoutes(api, handler)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", conf.App.Port)))
}
