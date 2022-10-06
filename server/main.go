package main

import (
	"fmt"
	"log"
	"os"

	_ "batleforc/bipper/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

// @title Bipper Api
// @version 1.0
// @description Bipper api

// @contact.name Batleforc
// @contact.url https://weebo.fr
// @contact.email maxleriche.60@gmail.com

// @BasePath /api
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
