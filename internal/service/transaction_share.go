package service

import (
	"context"
	"divvy/divvy-api/domain"
	"divvy/divvy-api/dto"

	"github.com/doug-martin/goqu/v9"
)

type TransactionShareService struct {
	transactionShareRepository domain.TransactionShareRepository
	db                         *goqu.Database
}

func NewTransactionShare(transactionShareRepository domain.TransactionShareRepository, db *goqu.Database) domain.TransactionShareService {
	return &TransactionShareService{
		transactionShareRepository: transactionShareRepository,
		db:                         db,
	}
}

// Create implements domain.TransactionShareService.
func (t *TransactionShareService) Create(ctx context.Context, req dto.CreateTransactionShareRequest) (dto.TransactionShareResponse, error) {
	panic("unimplemented")
}

// GetByTransaction implements domain.TransactionShareService.
func (t *TransactionShareService) GetByTransaction(ctx context.Context, id string) ([]dto.TransactionShareResponse, error) {
	panic("unimplemented")
}

// GetByUser implements domain.TransactionShareService.
func (t *TransactionShareService) GetByUser(ctx context.Context, uid string) ([]dto.TransactionShareResponse, error) {
	panic("unimplemented")
}

// Index implements domain.TransactionShareService.
func (t *TransactionShareService) Index(ctx context.Context) ([]dto.TransactionShareResponse, error) {
	panic("unimplemented")
}

// Show implements domain.TransactionShareService.
func (t *TransactionShareService) Show(ctx context.Context, id string) (dto.TransactionShareResponse, error) {
	panic("unimplemented")
}
