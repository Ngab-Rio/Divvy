package domain

import (
	"context"
	"database/sql"
	"divvy/divvy-api/dto"
	"time"
)

type Group struct {
	ID         string       `db:"id"`
	Name       string       `db:"name"`
	Created_by string       `db:"created_by"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}

type GroupWithUser struct {
	GroupID string `db:"group_id"`
	GroupName string `db:"group_name"`
	UserID string `db:"user_id"`
	Username string `db:"username"`
	Email string `db:"email"`
	Created_at sql.NullTime `db:"created_at"`
	Updated_at sql.NullTime `db:"updated_at"`
}

type GroupRepository interface{
	FindById(ctx context.Context, id string) (Group, error)
	FindByName(ctx context.Context, name string) (Group, error)
	GetAll(ctx context.Context) ([]Group, error)
	GetAllWithUser(ctx context.Context) ([]GroupWithUser, error)
	Save(ctx context.Context, g *Group) error
	// Delete(ctx context.Context, id string) error
}

type GroupService interface{
	Index(ctx context.Context) ([]dto.GroupResponse, error)
	IndexWithUser(ctx context.Context) ([]dto.GroupWithUserResponse, error)
	Create(ctx context.Context, req dto.CreateGroupRequest, userID string) (dto.GroupWithUserResponse ,error)
	// Delete(ctx context.Context, id string) error
}