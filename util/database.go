package util

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// dbProperties represents db configuration
type dbProperties struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func buildDBConfig(host string) *dbProperties {
	dbConfig := dbProperties{
		Host:     host,
		Port:     3306,
		User:     "user",
		Password: "password",
		DBName:   "wallet-tracky",
	}
	return &dbConfig
}

func getDbURL(dbConfig *dbProperties) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func  InitializeDatabase(config *viper.Viper){

	logger.Info("Initializing database.")

	env:=config.Get("env")

	var dbConfig *dbProperties

	if env=="dev"{
		dbConfig= buildDBConfig("localhost")
	}


	connectionURL:= getDbURL(dbConfig)

	db, err := gorm.Open(mysql.Open(connectionURL), &gorm.Config{})

	DB = db

	fmt.Println(DB)

	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("database initialized")
}
