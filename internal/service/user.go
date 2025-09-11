package service

import (
	"context"
	"divvy/divvy-api/domain"
	"divvy/divvy-api/dto"
	"errors"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &userService {
		userRepository: userRepository,
	}
}

func (u userService) Index(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := u.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var userData []dto.UserResponse
	for _, V := range users {
		userData = append(userData, dto.UserResponse{
			ID: V.ID,
			Username: V.Username,
			Email: V.Email,
		})
	}
	return userData, nil
}

func (u userService) Show(ctx context.Context, id string) (dto.UserResponse, error) {
	users, err := u.userRepository.FindById(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	if users.ID == "" {
		return dto.UserResponse{}, errors.New("users not found")
	}
	return dto.UserResponse{
		ID: users.ID,
		Username: users.Username,
		Email: users.Email,
	}, nil
}