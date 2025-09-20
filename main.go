package main

import (
	"divvy/divvy-api/internal/api"
	"divvy/divvy-api/internal/config"
	"divvy/divvy-api/internal/connection"
	"divvy/divvy-api/internal/middleware"
	"divvy/divvy-api/internal/repository"
	"divvy/divvy-api/internal/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()
	app.Use(middleware.CustomLogger())
	app.Use(middleware.CorsMiddleware())

	userRepository := repository.NewUser(dbConnection)
	groupRepository := repository.NewGroup(dbConnection)
	groupMemberRepository := repository.NewGroupMember(dbConnection)
	transactionRepository := repository.NewTransaction(dbConnection)

	authService := service.NewAuth(cnf, userRepository)
	userService := service.NewUser(userRepository)
	groupService := service.NewGroup(groupRepository, userRepository)
	groupMemberService := service.NewGroupMember(groupMemberRepository, groupRepository, userRepository)
	transactionService := service.NewTransaction(transactionRepository)

	api.NewAuth(app, authService)
	api.NewUser(app, userService, cnf.Jwt.Key)
	api.NewGroup(app, groupService, cnf.Jwt.Key)
	api.NewGroupMember(app, groupMemberService, cnf.Jwt.Key)
	api.NewTransaction(app, transactionService, cnf.Jwt.Key)

	addr := cnf.Server.Host + ":" + cnf.Server.Port
	fmt.Println("Server running at http://" + addr)

	if err := app.Listen(addr); err != nil {
		panic(err)
	}
}
