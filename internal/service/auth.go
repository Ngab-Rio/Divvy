package service

import (
	"context"
	"database/sql"
	"divvy/divvy-api/domain"
	"divvy/divvy-api/dto"
	"divvy/divvy-api/internal/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

func (a authService) Login (ctx context.Context, req dto.AuthLoginRequest) (dto.AuthResponse, error) {
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

func (a authService) Register (ctx context.Context, req dto.AuthRegisterRequest) (dto.AuthResponse, error) {
	user, _ := a.UserRepository.FindByEmail(ctx, req.Email)
	if user.Email != ""{
		return dto.AuthResponse{}, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.AuthResponse{}, errors.New("faild to hash password")
	}

	newUser := domain.User{
		ID: uuid.New().String(),
		Username: req.Username,
		Email: req.Email,
		Password: string(hashedPassword),
		Created_at: sql.NullTime{Valid: true, Time: time.Now()},
		Updated_at: sql.NullTime{Valid: true, Time: time.Now()},
	}

	err = a.UserRepository.Save(ctx, &newUser)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	claim := jwt.MapClaims{
		"id" : newUser.ID,
		"exp": time.Now().Add(time.Duration(a.conf.Jwt.Exp) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(a.conf.Jwt.Key))
	if err != nil {
		return dto.AuthResponse{}, errors.New("failed to generate token")
	}

	return dto.AuthResponse{
		Token: tokenStr,
	}, nil
}