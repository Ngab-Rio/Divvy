package dto

import (
	"time"
)

type GroupMemberResponse struct {
	ID                     string `json:"id"`
	GroupID                string `json:"group_id"`
	UserID                 string `json:"user_id"`
	Role                   string `json:"role"`
	JoinedAt time.Time `json:"joined_at"`
}

type BulkGroupMemberResponse struct {
	GroupID string `json:"group_id"`
	Members []GroupMemberResponse `json:"members"`
}

type CreateGroupMember struct {
	GroupID string `json:"group_id" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
	Role string `json:"role" validate:"required"`
}

// bulk insert
type CreateGroupMembersRequest struct {
	GroupID string `json:"group_id" validate:"required"`
	Members []CreateMemberPayload `json:"members" validate:"required,dive"`
}

type CreateMemberPayload struct {
	UserID string `json:"user_id" validate:"required"`
	Role string `json:"role" validate:"required"`
}