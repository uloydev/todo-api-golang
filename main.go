package main

import (
	"fmt"
	"todo-list-api/config"
	"todo-list-api/db"
	"todo-list-api/initialize"
	"todo-list-api/utils"

	_ "todo-list-api/docs"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title                       TODO LIST API
// @version                     1.0
// @description                 todo list api golang
// @contact.name                uloydev
// @contact.email               uloydev@gmail.com
// @BasePath                    /
func main() {
	conf := config.New()
	dbConn := db.NewGormConnection(conf)
	mailConf := config.NewMailConfig(conf)
	mailConn := utils.NewMailConnection(mailConf)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(compress.New())

	migrator := gormigrate.New(dbConn, gormigrate.DefaultOptions, db.Migrations)
	if err := migrator.Migrate(); err != nil {
		panic(err)
	}
	fmt.Println("SUCCESS migrate migrations")

	initialize.RunInitFunctions(app, dbConn, mailConn)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:          "http://localhost:3030/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	// use host:port/monitor to access monitor page
	app.Use(monitor.New(monitor.ConfigDefault))

	app.Listen(":3030")
}
