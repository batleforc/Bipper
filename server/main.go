package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "batleforc/bipper/docs"
	middle "batleforc/bipper/middleware"
	"batleforc/bipper/model"
	"batleforc/bipper/route"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Bipper Api
// @version 1.0
// @description Bipper api

// @contact.name Batleforc
// @contact.url https://weebo.fr
// @contact.email maxleriche.60@gmail.com

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @BasePath /api
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := model.ConnectDbFromEnv()
	model.InitDb(db)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${remote_ip} : ${time_rfc3339_nano}] ${status} : ${method} => ${uri}\n",
		Output: e.Logger.Output(),
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(os.Getenv("ALLOW_ORIGIN"), ","),
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderAccept},
	}))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})
	api := e.Group("/api")
	api.GET("/swagger/*", echoSwagger.WrapHandler)
	api.GET("/", func(c echo.Context) error {
		return c.String(200, "Hey I'm Bipper")
	})
	api.GET("/asset/:file", route.Asset)

	user := api.Group("/user")
	user.Use(middle.Auth())
	user.GET("", route.GetUser)
	user.POST("", route.SetUser)
	user.POST("/setpicture", route.SetPicture)

	auth := api.Group("/auth")
	auth.POST("/login", route.Login)
	auth.POST("/logout", route.Logout)
	auth.POST("/renew", route.RenewToken)
	auth.POST("/register", route.Register)

	channel := api.Group("/chan")
	channel.Use(middle.Auth())
	channel.POST("", route.CreateChan)
	channel.POST("/name", route.CheckChanName)
	channel.POST("/:chanId/renew", route.RenewChanPassword)
	channel.GET("/:chanId", route.GetOneChan)
	channel.GET("/:chanId/message", route.GetOneChanMessage)
	channel.GET("/public", route.GetPublicChannels)
	channel.GET("", route.GetUserChan)
	channel.DELETE("/:chanId", route.DeleteChan)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
