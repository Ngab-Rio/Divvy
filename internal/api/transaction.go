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

type transactionApi struct {
	transactionService domain.TransactionService
}

func NewTransaction(app *fiber.App, transactionService domain.TransactionService, secret string) {
	t := transactionApi{
		transactionService: transactionService,
	}

	transaction := app.Group("/transaction", middleware.JWTProtected(secret))

	transaction.Get("/", t.Index)
	transaction.Get("/:id", t.Show)
	transaction.Post("/", t.Create)
	transaction.Put("/:id", t.updateTransaction)
}

func (ta transactionApi) Index(ctx *fiber.Ctx) error{
	t, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	res, err := ta.transactionService.Index(t)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ta transactionApi) Create(ctx *fiber.Ctx) error {
	t, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req dto.CreateTransactionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)

	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	currentUserID := ctx.Locals("user_id").(string)

	res, err := ta.transactionService.Create(t, req, currentUserID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(res))
}

func(ta transactionApi) Show(ctx *fiber.Ctx) error {
	t, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")

	res, err := ta.transactionService.Show(t, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func(ta transactionApi) updateTransaction(ctx *fiber.Ctx) error {
	t, cancel := context.WithTimeout(ctx.Context(), time.Second * 10)
	defer cancel()

	var req dto.UpdateTransactionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)

	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	id := ctx.Params("id")


	res, err := ta.transactionService.Update(t, id, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(res))
}