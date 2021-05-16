package util

import (
	"fmt"
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

func buildDBConfig() *dbProperties {
	dbConfig := dbProperties{
		Host:     "localhost",
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

func  InitializeDatabase(){

	fmt.Println("Initializing database")

	dbConfig:= buildDBConfig()
	connectionURL:=getDbURL(dbConfig)

	db, err := gorm.Open(mysql.Open(connectionURL), &gorm.Config{})

	DB= db

	fmt.Println(DB)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("database initialized")
}
