package repository

import (
	"context"
	"database/sql"
	"divvy/divvy-api/domain"

	"github.com/doug-martin/goqu/v9"
)

type GroupRepository struct {
	db *goqu.Database
}

func NewGroup(con *sql.DB) domain.GroupRepository {
	return &GroupRepository{
		db: goqu.New("default", con),
	}
}

func (g GroupRepository) FindById(ctx context.Context, id string) (result domain.Group, err error){
	dataset := g.db.From("groups").Where(goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (g GroupRepository) FindByName(ctx context.Context, name string) (result domain.Group, err error){
	dataset := g.db.From("groups").Where(goqu.C("name").Eq(name))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (g GroupRepository) GetAll(ctx context.Context) (result []domain.Group, err error) {
	dataset := g.db.From("groups").Where(goqu.C("id").IsNotNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (g GroupRepository) GetAllWithUser(ctx context.Context) (result []domain.GroupWithUser, err error) {
	dataset := g.db.From("groups").Select(
		goqu.I("groups.id").As("group_id"),
		goqu.I("groups.name").As("group_name"),
		goqu.I("users.id").As("user_id"),
		goqu.I("users.username").As("username"),
		goqu.I("users.email").As("email"),
		goqu.I("groups.created_at"),
        goqu.I("groups.updated_at"),
	).Join(
		goqu.T("users"),
		goqu.On(goqu.I("groups.created_by").Eq(goqu.I("users.id"))),
	)

	err = dataset.ScanStructsContext(ctx, &result)
	return 
}

func (gr GroupRepository) Save(ctx context.Context, g *domain.Group) error {
	executor := gr.db.Insert("groups").Rows(g).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// func (gr GroupRepository) Delete(ctx context.Context, id string) error{
// 	executor := gr.db.Delete()
// }