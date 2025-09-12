package service

import (
	"context"
	"divvy/divvy-api/domain"
	"divvy/divvy-api/dto"
	"time"

	"github.com/google/uuid"
)

type groupMemberService struct {
	groupMemberRepository domain.GroupMemberRepository
	groupRepository domain.GroupRepository
	userRepository domain.UserRepository
}

func NewGroupMember(groupMemberRepository domain.GroupMemberRepository, groupRepository domain.GroupRepository, userRepository domain.UserRepository) domain.GroupMemberService {
	return &groupMemberService{
		groupMemberRepository: groupMemberRepository,
		groupRepository: groupRepository,
		userRepository: userRepository,
	}
}

func (gm *groupMemberService) Index(ctx context.Context) ([]dto.GroupMemberResponse, error) {
	groupMember, err := gm.groupMemberRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var groupMemberData []dto.GroupMemberResponse
	for _, v := range groupMember {
		groupMemberData = append(groupMemberData, dto.GroupMemberResponse{
			ID: v.ID,
			GroupID: v.GroupID,
			UserID: v.UserID,
			Role: string(v.Role),
			JoinedAt: v.JoinedAt,
		})
	}

	return groupMemberData, nil
}

func (gm groupMemberService) Create(ctx context.Context, req dto.CreateGroupMember) (dto.GroupMemberResponse, error) {
	groupMember := domain.GroupMember{
		ID: uuid.NewString(),
		GroupID: req.GroupID,
		UserID: req.UserID,
		Role: domain.MemberRole(req.Role),
		JoinedAt: time.Now(),
	}

	err := gm.groupMemberRepository.Save(ctx, &groupMember)
	if err != nil {
		return dto.GroupMemberResponse{}, err
	}

	resp := dto.GroupMemberResponse{
		ID: groupMember.ID,
		GroupID: groupMember.GroupID,
		UserID: groupMember.UserID,
		Role: string(groupMember.Role),
		JoinedAt: groupMember.JoinedAt,
	}
	return resp, nil
}

func (gm *groupMemberService) CreateBulk(ctx context.Context, req dto.CreateGroupMembersRequest) (dto.BulkGroupMemberResponse, error) {
	var members []dto.GroupMemberResponse

	tx, err := gm.groupMemberRepository.BeginTx(ctx)
	if err != nil {
		return  dto.BulkGroupMemberResponse{}, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	for _, m := range req.Members{
		gmData := domain.GroupMember{
			ID: uuid.New().String(),
			GroupID: req.GroupID,
			UserID: m.UserID,
			Role: domain.MemberRole(m.Role),
			JoinedAt: time.Now(),
		}

		if err = gm.groupMemberRepository.SaveTx(ctx, tx, &gmData); err != nil {
			return dto.BulkGroupMemberResponse{}, err
		}

		members = append(members, dto.GroupMemberResponse{
			ID: gmData.ID,
			GroupID: gmData.GroupID,
			UserID: gmData.UserID,
			Role: string(gmData.Role),
			JoinedAt: gmData.JoinedAt,
		})
	}

	if err = tx.Commit(); err != nil{
		return dto.BulkGroupMemberResponse{}, err
	}

	return dto.BulkGroupMemberResponse{
		GroupID: req.GroupID,
		Members: members,
	}, nil
}