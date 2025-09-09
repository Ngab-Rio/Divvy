package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/test", developer)

	app.Listen(":9000")

}

func developer(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")
}