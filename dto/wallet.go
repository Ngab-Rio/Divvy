package dto

import "time"

type WalletResponse struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Balance   float64 `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WalletWithTransactionsResponse struct {
	Wallet WalletResponse `json:"wallet"`
	Transactions []TransactionResponse `json:"transactions"`
}

type CreateWalletRequest struct {
	Balance float64 `json:"balance" validate:"gte=0"`
}