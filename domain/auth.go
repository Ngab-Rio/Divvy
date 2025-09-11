package domain

import (
	"context"
	"divvy/divvy-api/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthLoginRequest) (dto.AuthResponse, error)
	Register(ctx context.Context, req dto.AuthRegisterRequest) (dto.AuthResponse, error)
}