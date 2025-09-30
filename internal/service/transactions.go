package service

import (
	"context"
	"database/sql"
	"divvy/divvy-api/domain"
	"divvy/divvy-api/dto"
	"errors"
	"time"

	"github.com/google/uuid"
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
func (t *transactionService) Create(ctx context.Context, req dto.CreateTransactionRequest, currentUserID string) (dto.TransactionResponse, error) {
	if req.Amount <= 0 {
		return dto.TransactionResponse{}, errors.New("amount must be greater than zero")
	}

	tx := domain.Transaction{
		ID:          uuid.NewString(),
		GroupID:     sqlNullStringPtr(req.GroupID),
		CreatedBy:   currentUserID,
		PaidBy:      req.PaidBy,
		Amount:      req.Amount,
		Description: sqlNullString(req.Description),
		Date:        time.Now(),
		Type:        domain.TransactionType(req.Type),
		Source:      domain.TransactionSource(req.Source),
		ExternalRef: sqlNullString(req.ExternalRef),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := t.transactionRepository.Save(ctx, &tx); err != nil {
		return dto.TransactionResponse{}, err
	}

	return toTransactionResponse(tx), nil
}

// Delete implements domain.TransactionService.
func (t *transactionService) Delete(ctx context.Context, id string) error {
	existing, err := t.transactionRepository.FindByID(ctx, id)
	if err != nil {return err}
	if existing.ID == "" {return errors.New("transaction not found")}

	return  t.transactionRepository.Delete(ctx, id)
}

// GetByDateRange implements domain.TransactionService.
func (t *transactionService) GetByDateRange(ctx context.Context, groupID string, start time.Time, end time.Time) ([]dto.TransactionResponse, error) {
	panic("unimplemented")
}

// GetByGroup implements domain.TransactionService.
func (t *transactionService) GetByGroup(ctx context.Context, groupID string) ([]dto.TransactionResponse, error) {
	txs, err := t.transactionRepository.FindByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.TransactionResponse, 0, len(txs))
	for _, t := range txs {
		responses = append(responses, toTransactionWithDeatailResponse(t))
	}
	return  responses, nil
}

// GetBySource implements domain.TransactionService.
func (t *transactionService) GetBySource(ctx context.Context, groupID string, source domain.TransactionSource) ([]dto.TransactionResponse, error) {
	panic("unimplemented")
}

// GetByType implements domain.TransactionService.
func (t *transactionService) GetByType(ctx context.Context, groupID string, tType domain.TransactionType) ([]dto.TransactionResponse, error) {
	panic("unimplemented")
}

// Index implements domain.TransactionService.
func (t *transactionService) Index(ctx context.Context) ([]dto.TransactionResponse, error) {
	txs, err := t.transactionRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.TransactionResponse, 0, len(txs))
	for _, t := range txs {
		responses = append(responses, toTransactionWithDeatailResponse(t))
	}

	return responses, nil
}

// Show implements domain.TransactionService.
func (t *transactionService) Show(ctx context.Context, id string) (dto.TransactionResponse, error) {
	tx, err := t.transactionRepository.FindByID(ctx, id)
	if err != nil {
		return dto.TransactionResponse{}, err
	}
	return toTransactionWithDeatailResponse(tx), nil
}

// Update implements domain.TransactionService.
func (t *transactionService) Update(ctx context.Context, id string, req dto.UpdateTransactionRequest) (dto.TransactionResponse, error) {
	existing, err := t.transactionRepository.FindByIDRaw(ctx, id)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	if req.Amount != nil {
		existing.Amount = *req.Amount
	}

	if req.Description != nil {
		existing.Description = sqlNullString(*req.Description)
	}

	if req.Date != nil {
		existing.Date = *req.Date
	}

	if req.Type != nil {
		existing.Type = domain.TransactionType(*req.Type)
	}

	if req.Source != nil {
		existing.Source = domain.TransactionSource(*req.Source)
	}

	if req.ExternalRef != nil {
		existing.ExternalRef = sqlNullString(*req.ExternalRef)
	}

	existing.UpdatedAt = time.Now()

	if err := t.transactionRepository.Update(ctx, &existing); err != nil {
		return dto.TransactionResponse{}, err
	}

	return toTransactionResponse(existing), nil
}

func toTransactionResponse(t domain.Transaction) dto.TransactionResponse {
	var groupID *string
	if t.GroupID.Valid {
		groupID = &t.GroupID.String
	}

	return dto.TransactionResponse{
		ID:          t.ID,
		GroupID:     groupID,
		CreatedBy:   t.CreatedBy,
		PaidBy:      t.PaidBy,
		Amount:      t.Amount,
		Description: t.Description.String,
		Date:        t.Date,
		Type:        string(t.Type),
		Source:      string(t.Source),
		ExternalRef: t.ExternalRef.String,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
func toTransactionWithDeatailResponse(t domain.TransactionWithDetails) dto.TransactionResponse {
	var groupID *string
	if t.GroupID.Valid {
		groupID = &t.GroupID.String
	}

	return dto.TransactionResponse{
		ID:            t.ID,
		GroupID:       groupID,
		GroupName:     t.GroupName.String,
		CreatedBy:     t.CreatedBy,
		CreatedByName: t.CreatedByName,
		PaidBy:        t.PaidBy,
		PaidByName:    t.PaidByName,
		Amount:        t.Amount,
		Description:   t.Description.String,
		Date:          t.Date,
		Type:          string(t.Type),
		Source:        string(t.Source),
		ExternalRef:   t.ExternalRef.String,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
	}
}

// func mapTransactions(txs []domain.TransactionWithDetails) []dto.TransactionResponse {
// 	responses := make([]dto.TransactionResponse, 0, len(txs))
// 	for _, t := range txs {
// 		responses = append(responses, toTransactionWithDeatailResponse(t))
// 	}
// 	return responses
// }

func sqlNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func sqlNullStringPtr(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{}
}
