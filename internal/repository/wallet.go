package repository

import (
	"context"
	"database/sql"
	"divvy/divvy-api/domain"
	"errors"

	"github.com/doug-martin/goqu/v9"
)

type WalletRepository struct {
	db    *goqu.Database
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
func (w *WalletRepository) FindByUserID(ctx context.Context, uid string) (*domain.Wallet, error) {
	var wallet domain.Wallet
	dataset := w.db.From("wallets").Where(goqu.I("user_id").Eq(uid))
	found, err := dataset.ScanStructContext(ctx, &wallet)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, sql.ErrNoRows
	}

	return &wallet, nil
}

// Save implements domain.WalletRepository.
func (w *WalletRepository) Save(ctx context.Context, wallet *domain.Wallet) error {
	executor := w.db.Insert("wallets").Rows(wallet).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// Update implements domain.WalletRepository.
func (w *WalletRepository) Update(ctx context.Context, wallet *domain.Wallet) (domain.Wallet, error) {
	_, err := w.db.Update("wallets").Set(goqu.Record{
		"balance":    wallet.Balance,
		"updated_at": goqu.L("NOW()"),
	}).Where(goqu.C("user_id").Eq(wallet.User_id)).Executor().ExecContext(ctx)
	if err != nil {
		return domain.Wallet{}, err
	}

	var updated domain.Wallet
	_, err = w.db.From("wallets").Where(goqu.C("user_id").Eq(wallet.User_id)).ScanStructContext(ctx, &updated)
	return updated, err
}

// UpdateBalanceTx implements domain.WalletRepository.
func (w *WalletRepository) UpdateBalanceTx(ctx context.Context, sqlTx *sql.Tx, walletID string, delta float64) error {
	query, args, _ := w.db.Update("wallets").Set(goqu.Record{
		"balance" : goqu.L("balance + ?", delta),
		"updated_at" : goqu.L("NOW()"),
	}).Where(goqu.C("id").Eq(walletID), goqu.L("balance + ? >= 0", delta)).ToSQL()

	res, err := sqlTx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("wallet not found or insufficient balance")
	}

	return nil
}