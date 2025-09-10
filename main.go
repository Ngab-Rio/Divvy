package main

import (
	"divvy/divvy-api/internal/api"
	"divvy/divvy-api/internal/config"
	"divvy/divvy-api/internal/connection"
	"divvy/divvy-api/internal/repository"
	"divvy/divvy-api/internal/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	userRepository := repository.NewUser(dbConnection)

	authService := service.NewAuth(cnf, userRepository)

	api.NewAuth(app, authService)

	addr := cnf.Server.Host + ":" + cnf.Server.Port
	fmt.Println("Server running at http://" + addr)

	if err := app.Listen(addr); err != nil {
		panic(err)
	}
}
