package domain

import (
	"context"
	"database/sql"
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
}