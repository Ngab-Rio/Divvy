package main

import (
	"divvy/divvy-api/internal/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()

	app := fiber.New()

	addr := cnf.Server.Host + ":" + cnf.Server.Port
	fmt.Println("Server running at http://" + addr)

	if err := app.Listen(addr); err != nil {
		panic(err)
	}
}
