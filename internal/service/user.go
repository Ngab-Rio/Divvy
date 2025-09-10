package service

import (
	"context"
	"divvy/divvy-api/domain"
	"divvy/divvy-api/dto"
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