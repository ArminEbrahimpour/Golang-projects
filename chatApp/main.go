package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// ping handler
	app.Get("/ping", func(ctx *fiber.Ctx) error {

		return ctx.SendString("welcome to you ")

	})

	app.Listen(":7766")
}
