package domain

import (
	"context"
	"database/sql"
	"divvy/divvy-api/dto"
	"time"
)

type MemberRole string

const (
	RoleAdmin MemberRole = "admin"
	RoleMember MemberRole = "member"
)

type GroupMember struct {
	ID                     string  `db:"id"`
	GroupID                string  `db:"group_id"`
	UserID                 string  `db:"user_id"`
	Role                   MemberRole  `db:"role"`
	JoinedAt               time.Time `db:"joined_at"`
}

type GroupMemberWithMember struct {
	GroupMemberID 			string `db:"group_member_id"`
	GroupID 				string `db:"group_id"`
	GroupName 				string `db:"group_name"`
	UserID 					string `db:"user_id"`
	Username 				string `db:"username"`
	Email 					string `db:"email"`
	Role                   	MemberRole  `db:"role"`
	DefaultSharePercentage 	sql.NullFloat64 `db:"default_share_percentage"`
	JoinedAt               	time.Time `db:"joined_at"`
}

type GroupMemberRepository interface {
	FindById(ctx context.Context, id string) (GroupMember, error)
	FindByGroupID(ctx context.Context, groupID string) ([]GroupMember, error)
	GetAll(ctx context.Context) ([]GroupMember, error)
	GetAllWithMember(ctx context.Context) ([]GroupMemberWithMember, error)
	Save(ctx context.Context, gm *GroupMember) error
	SaveTx(ctx context.Context, tx *sql.Tx, gm *GroupMember) error
	BeginTx(ctx context.Context) (*sql.Tx, error)
}

type GroupMemberService interface{
	Index(ctx context.Context) ([]dto.GroupMemberResponse, error)
	Create(ctx context.Context, currentUserID string , req dto.CreateGroupMember) (dto.GroupMemberResponse, error)
	CreateBulk(ctx context.Context, req dto.CreateGroupMembersRequest) (dto.BulkGroupMemberResponse, error)
}