package api

import (
	"context"
	"divvy/divvy-api/domain"
	"divvy/divvy-api/dto"
	"divvy/divvy-api/internal/middleware"
	"divvy/divvy-api/internal/util"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type groupApi struct {
	groupService domain.GroupService
}

func NewGroup(app *fiber.App, groupService domain.GroupService, secret string){
	ga := groupApi{
		groupService: groupService,
	}

	groups := app.Group("/groups", middleware.JWTProtected(secret))

	groups.Get("/", ga.Index)
	groups.Get("/with-users", ga.IndexWithUser)
	groups.Post("/", ga.Create)
}

func (ga groupApi) Index(ctx *fiber.Ctx) error {
	g, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	res, err := ga.groupService.Index(g)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ga groupApi) IndexWithUser(ctx *fiber.Ctx) error {
	g, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	res, err := ga.groupService.IndexWithUser(g)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ga groupApi) Create(ctx *fiber.Ctx) error{
	g, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req dto.CreateGroupRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)

	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	userID := ctx.Locals("user_id").(string)

	res, err := ga.groupService.Create(g, req, userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(res))
}