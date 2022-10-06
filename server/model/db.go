package model

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SSL string

const (
	Enabled  SSL = "enable"
	Disabled SSL = "disable"
)

// https://gorm.io/docs/index.html
func InitDb(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Token{})
	db.AutoMigrate(&Channel{})
	db.AutoMigrate(&ChannelUser{})
	db.AutoMigrate(&Message{})
	return db
}

func ConnectDb(host, user, password, dbname, port string, ssl SSL) *gorm.DB {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, ssl)), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

func ConnectDbFromEnv() *gorm.DB {
	return ConnectDb(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), SSL(os.Getenv("DB_SSL")))
}
