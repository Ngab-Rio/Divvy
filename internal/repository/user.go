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

func (ur UserRepository) FindById(ctx context.Context, id string) (usr domain.User, err error) {
	dataset := ur.db.From("users").Where(goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &usr)
	return
}

func (ur UserRepository) GetAll(ctx context.Context) (result []domain.User, err error) {
	dataset := ur.db.From("users").Where(goqu.C("id").IsNotNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (ur UserRepository) Save(ctx context.Context, u *domain.User) error {
	executor := ur.db.Insert("users").Rows(u).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}