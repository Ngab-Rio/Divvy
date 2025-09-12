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

type groupMemberApi struct {
	groupMemberService domain.GroupMemberService
}

func NewGroupMember(app *fiber.App, groupMemberService domain.GroupMemberService, secret string) {
	gm := groupMemberApi{
		groupMemberService: groupMemberService,
	}

	groupsMember := app.Group("/group-members", middleware.JWTProtected(secret))

	groupsMember.Get("/", gm.Index)
	groupsMember.Post("/", gm.Create)
}

func (gmApi groupMemberApi) Index(ctx *fiber.Ctx) error{
	gm, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()
	
	res, err := gmApi.groupMemberService.Index(gm)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (gmApi groupMemberApi) Create(ctx *fiber.Ctx) error {
	gm, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req dto.CreateGroupMember
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)

	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	res, err := gmApi.groupMemberService.Create(gm, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(res))
}