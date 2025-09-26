package repository

import (
	"context"
	"database/sql"
	"divvy/divvy-api/domain"

	"github.com/doug-martin/goqu/v9"
)

type WalletRepository struct {
	db *goqu.Database
	sqlDB *sql.DB
}


func NewWallet(con *sql.DB) domain.WalletRepository {
	return &WalletRepository{
		db:    goqu.New("default", con),
		sqlDB: con,
	}
}

// Delete implements domain.WalletRepository.
func (w *WalletRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindById implements domain.WalletRepository.
func (w *WalletRepository) FindById(ctx context.Context, id string) (wallet domain.Wallet, err error) {
	dataset := w.db.From("wallets").Where(goqu.I("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &wallet)
	return
}

// FindByUserID implements domain.WalletRepository.
func (w *WalletRepository) FindByUserID(ctx context.Context, uid string) (wallet domain.Wallet, err error) {
	dataset := w.db.From("wallets").Where(goqu.I("user_id").Eq(uid))
	_, err = dataset.ScanStructContext(ctx, &wallet)
	return
}

// Save implements domain.WalletRepository.
func (w *WalletRepository) Save(ctx context.Context, wallet *domain.Wallet) error {
	executor := w.db.Insert("wallets").Rows(wallet).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// Update implements domain.WalletRepository.
func (w *WalletRepository) Update(ctx context.Context, wallet *domain.Wallet, uid string) ([]domain.Wallet, error) {
	panic("unimplemented")
}