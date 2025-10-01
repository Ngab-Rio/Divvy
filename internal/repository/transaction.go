package repository

import (
	"context"
	"database/sql"
	"divvy/divvy-api/domain"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type TransactionRepository struct {
	db    *goqu.Database
	sqlDB *sql.DB
}

// FindByIDRaw implements domain.TransactionRepository.

func NewTransaction(con *sql.DB) domain.TransactionRepository {
	return &TransactionRepository{
		db:    goqu.New("default", con),
		sqlDB: con,
	}
}

func (r *TransactionRepository) FindByIDRaw(ctx context.Context, id string) (tx domain.Transaction, err error) {
	dataset := r.db.From("transactions").Where(goqu.I("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &tx)
	return
}

func (r *TransactionRepository) FindByID(ctx context.Context, id string) (tx domain.TransactionWithDetails, err error) {
	dataset := r.db.From("transactions").
		Select(
			goqu.I("transactions.id"),
			goqu.I("transactions.group_id"),
			goqu.I("groups.name").As("group_name"),
			goqu.I("transactions.created_by"),
			goqu.I("creator.username").As("created_by_name"),
			goqu.I("transactions.paid_by"),
			goqu.I("payer.username").As("paid_by_name"),
			goqu.I("transactions.amount"),
			goqu.I("transactions.description"),
			goqu.I("transactions.date"),
			goqu.I("transactions.type"),
			goqu.I("transactions.source"),
			goqu.I("transactions.external_ref"),
			goqu.I("transactions.created_at"),
			goqu.I("transactions.updated_at"),
		).
		LeftJoin(goqu.T("groups"), goqu.On(goqu.Ex{"transactions.group_id": goqu.I("groups.id")})).
		Join(goqu.T("users").As("creator"), goqu.On(goqu.Ex{"transactions.created_by": goqu.I("creator.id")})).
		Join(goqu.T("users").As("payer"), goqu.On(goqu.Ex{"transactions.paid_by": goqu.I("payer.id")})).
		Where(goqu.Ex{"transactions.id": id})
	_, err = dataset.ScanStructContext(ctx, &tx)
	return tx, err
}

// FindByGroupID implements domain.TransactionRepository.
func (r *TransactionRepository) FindByGroupID(ctx context.Context, groupID string) (tx []domain.TransactionWithDetails, err error) {
	dataset := r.db.From("transactions").
		Select(
			goqu.I("transactions.id"),
			goqu.I("transactions.group_id"),
			goqu.I("groups.name").As("group_name"),
			goqu.I("transactions.created_by"),
			goqu.I("creator.username").As("created_by_name"),
			goqu.I("transactions.paid_by"),
			goqu.I("payer.username").As("paid_by_name"),
			goqu.I("transactions.amount"),
			goqu.I("transactions.description"),
			goqu.I("transactions.date"),
			goqu.I("transactions.type"),
			goqu.I("transactions.source"),
			goqu.I("transactions.external_ref"),
			goqu.I("transactions.created_at"),
			goqu.I("transactions.updated_at"),
		).
		LeftJoin(goqu.T("groups"), goqu.On(goqu.Ex{"transactions.group_id": goqu.I("groups.id")})).
		LeftJoin(goqu.T("users").As("creator"), goqu.On(goqu.Ex{"transactions.created_by": goqu.I("creator.id")})).
		LeftJoin(goqu.T("users").As("payer"), goqu.On(goqu.Ex{"transactions.paid_by": goqu.I("payer.id")})).
		Where(goqu.Ex{"transactions.group_id": groupID})
	sql, args, _ := dataset.ToSQL()
	fmt.Println("SQL QUERY:", sql)
	fmt.Println("ARGS:", args)

	err = dataset.ScanStructsContext(ctx, &tx)
	return
}

// GetAll implements domain.TransactionRepository.
func (r *TransactionRepository) GetAll(ctx context.Context) (txs []domain.TransactionWithDetails, err error) {
	dataset := r.db.From("transactions").
		Select(
			goqu.I("transactions.id"),
			goqu.I("transactions.group_id"),
			goqu.I("groups.name").As("group_name"),
			goqu.I("transactions.created_by"),
			goqu.I("creator.username").As("created_by_name"),
			goqu.I("transactions.paid_by"),
			goqu.I("payer.username").As("paid_by_name"),
			goqu.I("transactions.amount"),
			goqu.I("transactions.description"),
			goqu.I("transactions.date"),
			goqu.I("transactions.type"),
			goqu.I("transactions.source"),
			goqu.I("transactions.external_ref"),
			goqu.I("transactions.created_at"),
			goqu.I("transactions.updated_at"),
		).
		LeftJoin(goqu.T("groups"), goqu.On(goqu.Ex{"transactions.group_id": goqu.I("groups.id")})).
		LeftJoin(goqu.T("users").As("creator"), goqu.On(goqu.Ex{"transactions.created_by": goqu.I("creator.id")})).
		LeftJoin(goqu.T("users").As("payer"), goqu.On(goqu.Ex{"transactions.paid_by": goqu.I("payer.id")}))

	err = dataset.ScanStructsContext(ctx, &txs)
	return
}

// BeginTx implements domain.TransactionRepository.
func (r *TransactionRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return r.db.Db.BeginTx(ctx, nil)
}

// Delete implements domain.TransactionRepository.
func (r *TransactionRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Delete("transactions").Where(goqu.C("id").Eq(id)).Executor().ExecContext(ctx)
	return err
}

// FindByDateRange implements domain.TransactionRepository.
func (r *TransactionRepository) FindByDateRange(ctx context.Context, groupID string, start time.Time, end time.Time) (txs []domain.Transaction, err error) {
	dataset := r.db.From("transactions").Where(goqu.C("group_id").Eq(groupID), goqu.C("date").Between(goqu.Range(start, end)))
	err = dataset.ScanStructsContext(ctx, &txs)
	return
}

// FindBySource implements domain.TransactionRepository.
func (r *TransactionRepository) FindBySource(ctx context.Context, groupID string, souce domain.TransactionSource) (txs []domain.Transaction, err error) {
	dataset := r.db.From("transactions").Where(goqu.C("group_id").Eq(groupID), goqu.C("source").Eq(souce))
	err = dataset.ScanStructsContext(ctx, &txs)
	return
}

// FindByType implements domain.TransactionRepository.
func (r *TransactionRepository) FindByType(ctx context.Context, groupID string, tType domain.TransactionType) (txs []domain.Transaction, err error) {
	dataset := r.db.From("transactions").Where(goqu.C("group_id").Eq(groupID), goqu.C("type").Eq(tType))
	err = dataset.ScanStructsContext(ctx, &txs)
	return
}

// Save implements domain.TransactionRepository.
func (r *TransactionRepository) Save(ctx context.Context, tx *domain.Transaction) error {
	_, err := r.db.Insert("transactions").Rows(tx).Executor().ExecContext(ctx)
	return err
}

// SaveTx implements domain.TransactionRepository.
func (r *TransactionRepository) SaveTx(ctx context.Context, sqlTx *sql.Tx, tx *domain.Transaction) error {
	query, args, _ := r.db.Insert("transactions").Rows(tx).ToSQL()
	fmt.Println("QUERY:", query)
	fmt.Println("ARGS:", args)
	_, err := sqlTx.ExecContext(ctx, query, args...)
	return err
}

// Update implements domain.TransactionRepository.
func (r *TransactionRepository) Update(ctx context.Context, tx *domain.Transaction) error {
	_, err := r.db.Update("transactions").Set(goqu.Record{
		"amount":      tx.Amount,
		"description": tx.Description,
		"type":        tx.Type,
		"source":      tx.Source,
		"date":        tx.Date,
		"updated_at":  time.Now(),
	}).Where(goqu.C("id").Eq(tx.ID)).Executor().ExecContext(ctx)
	return err
}
