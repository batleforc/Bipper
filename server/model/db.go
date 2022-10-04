package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SSL string

const (
	Enabled  SSL = "enable"
	Disabled SSL = "disable"
)

func InitDb(host, user, password, dbname, port string, ssl SSL) *gorm.DB {

	return ConnectDb(host, user, password, dbname, port, ssl)
}

func ConnectDb(host, user, password, dbname, port string, ssl SSL) *gorm.DB {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, ssl)), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	return db
}
