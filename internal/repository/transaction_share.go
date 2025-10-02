package repository

import (
	"context"
	"database/sql"
	"divvy/divvy-api/domain"

	"github.com/doug-martin/goqu/v9"
)

type TransactionSHareRepository struct {
	db    *goqu.Database
	sqlDB *sql.DB
}


func NewTransactionShare(con *sql.DB) domain.TransactionShareRepository {
	return &TransactionSHareRepository{
		db:    goqu.New("default", con),
		sqlDB: con,
	}
}

// Delete implements domain.TransactionShareRepository.
func (t *TransactionSHareRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindByID implements domain.TransactionShareRepository.
func (t *TransactionSHareRepository) FindByID(ctx context.Context, id string) (domain.TransactionShare, error) {
	panic("unimplemented")
}

// FindByStatus implements domain.TransactionShareRepository.
func (t *TransactionSHareRepository) FindByStatus(ctx context.Context, uid string, status domain.ShareTransactionStatus) ([]domain.TransactionShare, error) {
	panic("unimplemented")
}

// FindByTransactionID implements domain.TransactionShareRepository.
func (t *TransactionSHareRepository) FindByTransactionID(ctx context.Context, id string) (domain.TransactionShare, error) {
	panic("unimplemented")
}

// FindByUserID implements domain.TransactionShareRepository.
func (t *TransactionSHareRepository) FindByUserID(ctx context.Context, id string) (domain.TransactionShare, error) {
	panic("unimplemented")
}

// GetAll implements domain.TransactionShareRepository.
func (t *TransactionSHareRepository) GetAll(ctx context.Context) ([]domain.TransactionShare, error) {
	panic("unimplemented")
}

// Save implements domain.TransactionShareRepository.
func (t *TransactionSHareRepository) Save(ctx context.Context, tx *domain.TransactionShare) error {
	panic("unimplemented")
}

// Update implements domain.TransactionShareRepository.
func (t *TransactionSHareRepository) Update(ctx context.Context, tx *domain.TransactionShare) error {
	panic("unimplemented")
}