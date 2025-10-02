package dto

type CreateTransactionShareRequest struct {
	TransactionID string `json:"transaction_id" validate:"required"`
	UserID        string `json:"user_id" validate:"required"`
	ShareAmount float64 `json:"share_amount" validate:"required"`
	Status string `json:"status" validate:"required,oneof=pending paid"`
}

type TransactionShareResponse struct {
	ID string `json:"id"`
	TransactionID string `json:"transaction_id"`
	Description string `json:"description,omitempty"`
	UserID string `json:"user_id"`
	Username string `json:"username"`
	ShareAmount float64 `json:"share_amount"`
	Status string `json:"status"`
}
