package main

import (
	"fmt"

	config "github.com/husfuu/crowdfunding-octo/config"
	"github.com/husfuu/crowdfunding-octo/internal/server"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/db"
	"github.com/husfuu/crowdfunding-octo/pkg/infra/logger"
)

func main() {
	conf := config.InitConf("dev")

	appLoger := logger.NewLogrusLogger(&conf.Logger.Logrus)

	var dbList db.DatabaseList
	dbList.APIDB = db.NewGormConnection(&conf.Connection.DatabaseApp, appLoger)
	fmt.Println("main: config: ", conf.Connection.DatabaseApp)
	server.Run(conf, &dbList, appLoger)
}
