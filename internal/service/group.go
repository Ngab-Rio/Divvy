package service

import (
	"context"
	"database/sql"
	"divvy/divvy-api/domain"
	"divvy/divvy-api/dto"
	"errors"
	"time"

	"github.com/google/uuid"
)

type groupService struct {
	groupRepository domain.GroupRepository
	userRepository domain.UserRepository
}

func NewGroup(groupRepository domain.GroupRepository, userRepository domain.UserRepository) domain.GroupService {
	return &groupService{
		groupRepository: groupRepository,
		userRepository: userRepository,
	}
}

func (g groupService) Index(ctx context.Context) ([]dto.GroupResponse, error) {
	groups, err := g.groupRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var groupData []dto.GroupResponse
	for _, v := range groups {
		groupData = append(groupData, dto.GroupResponse{
			ID: v.ID,
			Name: v.Name,
			Created_by: v.Created_by,
			Created_at: v.Created_at,
			Updated_at: v.Updated_at,
		})
	}
	return  groupData, nil
}

func (g groupService) IndexWithUser(ctx context.Context) ([]dto.GroupWithUserResponse, error) {
	groupsWithUser, err := g.groupRepository.GetAllWithUser(ctx)
	if err != nil {
		return nil, err
	}
	var groupData []dto.GroupWithUserResponse
	for _, v := range groupsWithUser {
		groupData = append(groupData, dto.GroupWithUserResponse{
			ID: v.GroupID,
			Name: v.GroupName,
			Created_by: dto.UserResponse{
				ID: v.UserID,
				Username: v.Username,
				Email: v.Email,
			},
			Created_at: v.Created_at.Time,
			Updated_at: v.Created_at.Time,
		})
	}

	return groupData, nil
}

func (g groupService) Create(ctx context.Context, req dto.CreateGroupRequest, userID string) (dto.GroupWithUserResponse,error) {
	existing, err := g.groupRepository.FindByName(ctx, req.Name)
	if err != nil && err != sql.ErrNoRows{
		return dto.GroupWithUserResponse{}, err
	}

	if existing.ID != "" {
		return dto.GroupWithUserResponse{}, errors.New("group name already exists")
	}
	
	group := domain.Group{
		ID: uuid.NewString(),
		Name: req.Name,
		Created_by: userID,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	
	if err := g.groupRepository.Save(ctx, &group); err != nil {
		return dto.GroupWithUserResponse{}, err
	}

	user, err := g.userRepository.FindById(ctx, userID)
	if err != nil {
		return dto.GroupWithUserResponse{}, err
	}

	return dto.GroupWithUserResponse{
		ID: group.ID,
		Name: group.Name,
		Created_by: dto.UserResponse{
			ID: user.ID,
			Username: user.Username,
			Email: user.Email,
		},
		Created_at: group.Created_at,
		Updated_at: group.Created_at,
	}, nil
}

// func (g groupService) Delete(ctx context.Context, id string) error {
// 	exists, err := g.groupRepository.FindById(ctx, id)
// 	if err != nil {return  err}
// 	if exists.ID == "" {return errors.New("group not found")}
// 	return g.groupRepository.Delete(ctx, id)
// }