package main

import (
	"todo-list-api/config"
	"todo-list-api/db"
	"todo-list-api/initialize"
	"todo-list-api/utils"

	_ "todo-list-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title                       TODO LIST API
// @version                     1.0
// @description                 todo list api golang
// @contact.name                uloydev
// @contact.email               wahyumiftahul7@gmail.com
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @BasePath                    /
func main() {
	conf := config.New()
	dbConn := db.NewGormConnection(conf)
	mailConf := config.NewMailConfig(conf)
	mailConn := utils.NewMailConnection(mailConf)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	v1 := app.Group("/v1")

	initialize.RunInitFunctions(v1.(*fiber.Group), dbConn, mailConn)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:          "http://localhost:6991/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	app.Listen(":6991")
}
