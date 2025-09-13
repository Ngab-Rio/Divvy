package repository

import (
	"context"
	"database/sql"
	"divvy/divvy-api/domain"

	"github.com/doug-martin/goqu/v9"
)

type GroupMemberRepository struct {
	db *goqu.Database
	sqlDB *sql.DB
}

func NewGroupMember(con *sql.DB) domain.GroupMemberRepository {
	return &GroupMemberRepository{
		db: goqu.New("default", con),
		sqlDB: con,
	}
}

func (gm GroupMemberRepository) FindById(ctx context.Context, id string) (result domain.GroupMember, err error) {
	dataset := gm.db.From("group_members").Where(goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (gm GroupMemberRepository) FindByGroupID(ctx context.Context, groupID string) (result []domain.GroupMemberWithMember, err error) {
	dataset := gm.db.From("group_members").Select(
		goqu.I("group_members.id").As("group_member_id"),
		goqu.I("group_members.group_id"),
		goqu.I("groups.name").As("group_name"),
		goqu.I("users.id").As("user_id"),
		goqu.I("users.username"),
		goqu.I("users.email"),
		goqu.I("group_members.role"),
		goqu.I("group_members.joined_at"),
	).Join(
		goqu.T("users"),
		goqu.On(goqu.I("group_members.user_id").Eq(goqu.I("users.id"))),
	).Join(
		goqu.T("groups"),
		goqu.On(goqu.I("group_members.group_id").Eq(goqu.I("groups.id"))),
	).Where(goqu.Ex{"group_members.group_id":groupID})
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (gm GroupMemberRepository) GetAll(ctx context.Context) (result []domain.GroupMember, err error) {
	dataset := gm.db.From("group_members")
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (gm GroupMemberRepository) GetAllWithMember(ctx context.Context) (result []domain.GroupMemberWithMember, err error) {
	dataset := gm.db.From("group_members").Select(
		goqu.I("group_members.id").As("group_member_id"),
		goqu.I("group_members.group_id"),
		goqu.I("groups.name").As("group_name"),
		goqu.I("users.id").As("user_id"),
		goqu.I("users.username"),
		goqu.I("users.email"),
		goqu.I("group_members.role"),
		goqu.I("group_members.joined_at"),
	).Join(goqu.T("groups"), goqu.On(goqu.Ex{"group_members.group_id": goqu.I("groups.id")})).Join(goqu.I("users"), goqu.On(goqu.Ex{"group_members.user_id": goqu.I("users.id")}))
	
	sqlStr, args, _ := dataset.ToSQL()
	err = gm.db.ScanStructsContext(ctx, &result, sqlStr, args...)
	return
}

func (gm GroupMemberRepository) Save(ctx context.Context, g *domain.GroupMember) error {
	executor := gm.db.Insert("group_members").Rows(g).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (gm *GroupMemberRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return gm.sqlDB.BeginTx(ctx, nil)
}

func (gm *GroupMemberRepository) SaveTx(ctx context.Context, tx *sql.Tx, g *domain.GroupMember) error {
	query, args, err := gm.db.Insert("group_members").Rows(goqu.Record{
		"id": g.ID,
		"group_id": g.GroupID,
		"user_id": g.UserID,
		"role": g.Role,
		"joined_at": g.JoinedAt,
	}).ToSQL()
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, query, args...)
	return err
}