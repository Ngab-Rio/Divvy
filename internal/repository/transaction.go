package repository

import (
	"context"
	"database/sql"
	"divvy/divvy-api/domain"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type TransactionRepository struct {
	db *goqu.Database
	sqlDB *sql.DB
}

func NewTransaction(con *sql.DB) domain.TransactionRepository {
	return &TransactionRepository{
		db:    goqu.New("default", con),
		sqlDB: con,
	}
}

func (r *TransactionRepository) FindByID(ctx context.Context, id string) (domain.Transaction, error) {
	var tx domain.Transaction
	dataset := r.db.From("transactions").Where(goqu.C("id").Eq(id))
	_, err := dataset.ScanStructContext(ctx, &tx)
	return tx, err
	
}

// FindByGroupID implements domain.TransactionRepository.
func (r *TransactionRepository) FindByGroupID(ctx context.Context, groupID string) (tx []domain.Transaction, err error) {
	dataset := r.db.From("transactions").Where(goqu.C("group_id").Eq(groupID))
	err = dataset.ScanStructsContext(ctx, &tx)
	return
}

// GetAll implements domain.TransactionRepository.
func (r *TransactionRepository) GetAll(ctx context.Context) (txs []domain.Transaction,err error) {
	dataset := r.db.From("transactions")
	err = dataset.ScanStructsContext(ctx, &txs)
	return
}

// BeginTx implements domain.TransactionRepository.
func (r *TransactionRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	panic("unimplemented")
}

// Delete implements domain.TransactionRepository.
func (r *TransactionRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindByDateRange implements domain.TransactionRepository.
func (r *TransactionRepository) FindByDateRange(ctx context.Context, groupID string, start time.Time, end time.Time) ([]domain.Transaction, error) {
	panic("unimplemented")
}


// FindBySource implements domain.TransactionRepository.
func (r *TransactionRepository) FindBySource(ctx context.Context, groupID string, souce domain.TransactionSource) ([]domain.Transaction, error) {
	panic("unimplemented")
}

// FindByType implements domain.TransactionRepository.
func (r *TransactionRepository) FindByType(ctx context.Context, groupID string, tType domain.TransactionType) ([]domain.Transaction, error) {
	panic("unimplemented")
}


// Save implements domain.TransactionRepository.
func (r *TransactionRepository) Save(ctx context.Context, tx *domain.Transaction) error {
	_, err := r.db.Insert("transactions").Rows(tx).Executor().ExecContext(ctx)
	return err
}

// SaveTx implements domain.TransactionRepository.
func (r *TransactionRepository) SaveTx(ctx context.Context, sqlTx *sql.Tx, tx *domain.Transaction) error {
	panic("unimplemented")
}

// Update implements domain.TransactionRepository.
func (r *TransactionRepository) Update(ctx context.Context, tx *domain.Transaction) error {
	panic("unimplemented")
}