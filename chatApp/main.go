package main

import (
	"chatApp/src/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	// create views engine
	viewsEngine := html.New("./views", ".html")

	// start app fiber
	app := fiber.New(fiber.Config{

		Views: viewsEngine,
	})

	// static route and dir
	app.Static("./static/", "./static")

	// ping handler
	app.Get("/ping", func(ctx *fiber.Ctx) error {

		return ctx.SendString("welcome to you ")

	})
	// creating new handler
	appHandler := handlers.NewAppHandler()

	/*
		routers:
	*/
	// app handler routes
	app.Get("/", appHandler.HandleGetIndex)

	app.Listen(":7766")
}
