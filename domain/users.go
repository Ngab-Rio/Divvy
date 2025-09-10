package domain

import (
	"context"
	"database/sql"
)

type Users struct {
	ID         string `db:"id"`
	Username   string `db:"username"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	Created_at sql.NullTime `db:"created_at"`
	Updated_at sql.NullTime `db:"updated_at"`
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]Users, error)
	FindByAll(ctx context.Context, id string) (Users, error)
	Save(ctx context.Context, u *Users) error
	Update(ctx context.Context, u *Users) error
	Delete(ctx context.Context, u *Users) error
}