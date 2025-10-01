package api

import (
	"context"
	"divvy/divvy-api/domain"
	"divvy/divvy-api/dto"
	"divvy/divvy-api/internal/middleware"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type walletAPI struct {
	walletService domain.WalletService
}

func NewWallet(app *fiber.App, walletService domain.WalletService, secret string) {
	wa := walletAPI{
		walletService: walletService,
	}

	wallets := app.Group("/wallets", middleware.JWTProtected(secret))

	wallets.Get("/:id", wa.getByID)
	wallets.Post("/", wa.CreateWallet)
	wallets.Put("/:id", wa.UpdateWallet)
}

func (wa walletAPI) getByID(ctx *fiber.Ctx) error {
	wallet, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	res, err := wa.walletService.GetWalletByID(wallet, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(res))
}

func (wa walletAPI) CreateWallet(ctx *fiber.Ctx) error {
	w, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req dto.CreateWalletRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError("invalid request body"))
	}

	currentID := ctx.Locals("user_id").(string)

	wallet, err := wa.walletService.CreateWallet(w, currentID, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(wallet))
}

func (wa walletAPI) UpdateWallet(ctx *fiber.Ctx) error {
	w, cancel := context.WithTimeout(ctx.Context(), time.Second * 10)
	defer cancel()

	var req dto.UpdateWalletRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError("invalid request body"))
	}

	id := ctx.Params("id")

	wallet, err := wa.walletService.UpdateWallet(w, id, req)
	if err != nil{
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(wallet))
}