package repository

import (
	"context"
	"database/sql"
	"divvy/divvy-api/domain"

	"github.com/doug-martin/goqu/v9"
)

type UserRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &UserRepository{
		db: goqu.New("default", con),
	}
}

func (u UserRepository) FindByEmail(ctx context.Context, email string) (usr domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.C("email").Eq(email))
	_, err = dataset.ScanStructContext(ctx, &usr)
	return
}

func (u UserRepository) GetAll(ctx context.Context) (result []domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.C("id").IsNotNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}