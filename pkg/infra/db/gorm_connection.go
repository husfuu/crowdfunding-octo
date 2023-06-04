package db

import (
	"fmt"

	"github.com/husfuu/crowdfunding-octo/config"
	"github.com/husfuu/crowdfunding-octo/pkg/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

type DatabaseList struct {
	APIDB *gorm.DB
}

func NewGormConnection(conf *config.DatabaseAccount, log *logrus.Logger) *gorm.DB {
	var db *gorm.DB
	var err error

	dbName := utils.GetDBNameFromDriverSource(conf.DriverSource)
	fmt.Println("dbName: ", dbName)
	if conf.DriverName == "postgres" || conf.DriverName == "pgx" {
		db, err = gorm.Open(postgres.Open(conf.DriverSource), &gorm.Config{
			Logger: gormlog.Default.LogMode(gormlog.LogLevel(gormlog.Error)),
		})
	}

	if err != nil {
		log.Fatal("Failed to connect database " + dbName + ", err: " + err.Error())
	}

	log.Info("Connection Opened to Database " + dbName)
	return db
}
