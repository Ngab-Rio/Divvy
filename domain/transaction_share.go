package domain

import (
	"context"
	"divvy/divvy-api/dto"
)

type ShareTransactionStatus string

const (
	Pending ShareTransactionStatus = "pending"
	Paid    ShareTransactionStatus = "paid"
)

type TransactionShare struct {
	ID                string  `db:"id"`
	TransactionID     string  `db:"transaction_id"`
	UserID            string  `db:"user_id"`
	ShareAmount       float64 `db:"share_amount"`
	StatusTransaction string  `db:"status"`
}

type TransactionShareRepository interface {
	FindByID(ctx context.Context, id string) (TransactionShare, error)
	FindByTransactionID(ctx context.Context, id string) (TransactionShare, error)
	FindByUserID(ctx context.Context, id string) (TransactionShare, error)
	GetAll(ctx context.Context) ([]TransactionShare, error)
	Save(ctx context.Context, tx *TransactionShare) error
	Update(ctx context.Context, tx *TransactionShare) error
	Delete(ctx context.Context, id string) error

	FindByStatus(ctx context.Context, uid string, status ShareTransactionStatus) ([]TransactionShare, error)
}

type TransactionShareService interface {
	Index(ctx context.Context) ([]dto.TransactionShareResponse, error)
	Show(ctx context.Context, id string) (dto.TransactionShareResponse, error)
	Create(ctx context.Context, req dto.CreateTransactionShareRequest) (dto.TransactionShareResponse, error)
	// Update(ctx context.Context, id strin)

	GetByUser(ctx context.Context, uid string) ([]dto.TransactionShareResponse, error)
	GetByTransaction(ctx context.Context, id string) ([]dto.TransactionShareResponse, error)
}