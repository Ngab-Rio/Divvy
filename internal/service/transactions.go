package service

import (
	"context"
	"divvy/divvy-api/domain"
	"divvy/divvy-api/dto"
	"time"
)

type transactionService struct {
	transactionRepository domain.TransactionRepository
}


func NewTransaction(transactionRepository domain.TransactionRepository) domain.TransactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
	}
}

// CalculateGroupBalance implements domain.TransactionService.
func (t *transactionService) CalculateGroupBalance(ctx context.Context, groupID string) (float64, error) {
	panic("unimplemented")
}

// Create implements domain.TransactionService.
func (t *transactionService) Create(ctx context.Context, req domain.Transaction) (domain.Transaction, error) {
	panic("unimplemented")
}

// Delete implements domain.TransactionService.
func (t *transactionService) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetByDateRange implements domain.TransactionService.
func (t *transactionService) GetByDateRange(ctx context.Context, groupID string, start time.Time, end time.Time) ([]domain.Transaction, error) {
	panic("unimplemented")
}

// GetByGroup implements domain.TransactionService.
func (t *transactionService) GetByGroup(ctx context.Context, groupID string) ([]domain.Transaction, error) {
	panic("unimplemented")
}

// GetBySource implements domain.TransactionService.
func (t *transactionService) GetBySource(ctx context.Context, groupID string, source domain.TransactionSource) ([]domain.Transaction, error) {
	panic("unimplemented")
}

// GetByType implements domain.TransactionService.
func (t *transactionService) GetByType(ctx context.Context, groupID string, tType domain.TransactionType) ([]domain.Transaction, error) {
	panic("unimplemented")
}

// Index implements domain.TransactionService.
func (t *transactionService) Index(ctx context.Context) ([]domain.Transaction, error) {
	// txs, err := t.transactionRepository.GetAll(ctx)
	// if err != nil{
	// 	return nil, err
	// }

	// responses := make([]dto.TransactionResponse, 0, len(txs))
	// for _, t := range txs{
	// 	responses = append(responses, toTransactionResponse(t))
	// }

	// return responses, nil
	panic("unimplemented")
}

// Show implements domain.TransactionService.
func (t *transactionService) Show(ctx context.Context, id string) (domain.Transaction, error) {
	panic("unimplemented")
}

// Update implements domain.TransactionService.
func (t *transactionService) Update(ctx context.Context, id string, req domain.Transaction) (domain.Transaction, error) {
	panic("unimplemented")
}

func toTransactionResponse(t domain.Transaction) dto.TransactionResponse {
	return dto.TransactionResponse{
		ID: t.ID,
		GroupID: t.GroupID,
		CreatedBy: t.CreatedBy,
		PaidBy: t.PaidBy,
		Amount: t.Amount,
		Description: t.Description.String,
		Date: t.Date,
		Type: string(t.Type),
		Source: string(t.Source),
		ExternalRef: t.ExternalRef.String,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}