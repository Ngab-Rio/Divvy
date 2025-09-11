package domain

import (
	"context"
	"database/sql"
	"divvy/divvy-api/dto"
)

type User struct {
	ID         string `db:"id"`
	Username   string `db:"username"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	Created_at sql.NullTime `db:"created_at"`
	Updated_at sql.NullTime `db:"updated_at"`
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (User, error)
	FindById(ctx context.Context, id string) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	Save(ctx context.Context, u *User) error
}

type UserService interface {
	Index(ctx context.Context) ([]dto.UserResponse, error)
	Show(ctx context.Context, id string) (dto.UserResponse, error)
}