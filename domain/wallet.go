package domain

import (
	"context"
	"database/sql"
	"divvy/divvy-api/dto"
	"time"
)

type Wallet struct {
	ID        string  `db:"id"`
	User_id   string  `db:"user_id"`
	Balance   float64 `db:"balance"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type WalletRepository interface {
	FindById(ctx context.Context, id string) (Wallet, error)
	FindByUserID(ctx context.Context, uid string) (*Wallet, error)
	Save(ctx context.Context, wallet *Wallet) error
	Update(ctx context.Context, wallet *Wallet) (Wallet, error)
	Delete(ctx context.Context, id string) error
	UpdateBalanceTx(ctx context.Context, sqlTx *sql.Tx, walletID string, delta float64) error
}

type WalletService interface {
	CreateWallet(ctx context.Context, userID string, req dto.CreateWalletRequest) (dto.WalletResponse, error)
	GetWalletByID(ctx context.Context, id string) (dto.WalletResponse, error)
	GetWalletByUser(ctx context.Context, uid string) ([]dto.WalletResponse, error)
	UpdateWallet(ctx context.Context, id string, req dto.UpdateWalletRequest) (dto.WalletResponse, error)
	DeleteWallet(ctx context.Context, id string) error

	GetWalletWithTransactions(ctx context.Context, walletID string) (dto.WalletWithTransactionsResponse, error)
	GetWalletsWithTransactions(ctx context.Context, userID string) ([]dto.WalletWithTransactionsResponse, error)
}