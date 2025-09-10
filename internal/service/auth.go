package service

import (
	"context"
	"divvy/divvy-api/domain"
	"divvy/divvy-api/dto"
	"divvy/divvy-api/internal/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	conf *config.Config
	UserRepository domain.UserRepository
}

func NewAuth (cnf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return authService{
		conf: cnf,
		UserRepository: userRepository,
	}
}

func (a authService) Login (ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := a.UserRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	if user.ID == "" {
		return dto.AuthResponse{}, errors.New("id user kosong")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}

	claim := jwt.MapClaims{
		"id": user.ID,
		"exp": time.Now().Add(time.Duration(a.conf.Jwt.Exp) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenStr, err := token.SignedString([]byte(a.conf.Jwt.Key))
	if err != nil {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}

	return dto.AuthResponse{
		Token: tokenStr,
	}, nil
}