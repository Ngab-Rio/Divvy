package dto

import "time"

type CreateTransactionRequest struct {
	GroupID     *string   `json:"group_id,omitempty"`
	CreatedBy   string    `json:"created_by" validate:"required"`
	PaidBy      string    `json:"paid_by" validate:"required"`
	Amount      float64   `json:"amount" validate:"required"`
	Description string    `json:"description,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	Type        string    `json:"type" validate:"required,oneof=income expense"`
	Source      string    `json:"source" validate:"required,oneof=manual gmail_brimo import_csv qr_scan"`
	ExternalRef string    `json:"external_ref,omitempty"`
}

type UpdateTransactionRequest struct {
	// ID          string    `json:"id" validate:"required"`
	Amount      *float64   `json:"amount,omitempty"`
	Description *string    `json:"description,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
	Type        *string    `json:"type,omitempty" validate:"omitempty,oneof=income expense"`
	Source      *string    `json:"source,omitempty" validate:"omitempty,oneof=manual gmail_brimo import_csv qr_scan"`
	ExternalRef *string    `json:"external_ref,omitempty"`
}

type TransactionResponse struct {
	ID            string    `json:"id"`
	GroupID       *string   `json:"group_id,omitempty"`
	GroupName     string    `json:"group_name,omitempty"`
	CreatedBy     string    `json:"created_by"`
	CreatedByName string    `json:"created_by_name,omitempty"`
	PaidBy        string    `json:"paid_by"`
	PaidByName    string    `json:"paid_by_name,omitempty"`
	Amount        float64   `json:"amount"`
	Description   string    `json:"description,omitempty"`
	Date          time.Time `json:"date"`
	Type          string    `json:"type"`
	Source        string    `json:"source"`
	ExternalRef   string    `json:"external_ref,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type TranscationListResponse struct {
	GroupID      string                `json:"group_id"`
	Transactions []TransactionResponse `json:"transactions"`
	TotalAmount  float64               `json:"total_amount,omitempty"`
	Balance      float64               `json:"balance,omitempty"`
}
