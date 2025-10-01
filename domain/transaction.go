package domain

import (
	"context"
	"database/sql"
	"divvy/divvy-api/dto"
	"time"
)

type TransactionType string
type TransactionSource string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"

	Manual     TransactionSource = "manual"
	GmailBrimo TransactionSource = "gmail_brimo"
	ImportCSV  TransactionSource = "import_csv"
	QRScan     TransactionSource = "qr_scan"
)

type Transaction struct {
	ID          string            `db:"id"`
	GroupID     sql.NullString    `db:"group_id"`
	WalletID    string            `db:"wallet_id"`
	CreatedBy   string            `db:"created_by"`
	PaidBy      string            `db:"paid_by"`
	Amount      float64           `db:"amount"`
	Description sql.NullString    `db:"description"`
	Date        time.Time         `db:"date"`
	Type        TransactionType   `db:"type"`
	Source      TransactionSource `db:"source"`
	ExternalRef sql.NullString    `db:"external_ref"`
	CreatedAt   time.Time         `db:"created_at"`
	UpdatedAt   time.Time         `db:"updated_at"`
}

type TransactionWithDetails struct {
	ID            string            `db:"id"`
	GroupID       sql.NullString    `db:"group_id"`
	WalletID      sql.NullString    `db:"wallet_id"`
	GroupName     sql.NullString    `db:"group_name"`
	CreatedBy     string            `db:"created_by"`
	CreatedByName string            `db:"created_by_name"`
	PaidBy        string            `db:"paid_by"`
	PaidByName    string            `db:"paid_by_name"`
	Amount        float64           `db:"amount"`
	Description   sql.NullString    `db:"description"`
	Date          time.Time         `db:"date"`
	Type          TransactionType   `db:"type"`
	Source        TransactionSource `db:"source"`
	ExternalRef   sql.NullString    `db:"external_ref"`
	CreatedAt     time.Time         `db:"created_at"`
	UpdatedAt     time.Time         `db:"updated_at"`
}

type TransactionRepository interface {
	FindByID(ctx context.Context, id string) (TransactionWithDetails, error)
	FindByIDRaw(ctx context.Context, id string) (Transaction, error)
	FindByGroupID(ctx context.Context, groupID string) ([]TransactionWithDetails, error)
	GetAll(ctx context.Context) ([]TransactionWithDetails, error)
	Save(ctx context.Context, tx *Transaction) error
	Update(ctx context.Context, tx *Transaction) error
	Delete(ctx context.Context, id string) error

	SaveTx(ctx context.Context, sqlTx *sql.Tx, tx *Transaction) error
	BeginTx(ctx context.Context) (*sql.Tx, error)

	FindByDateRange(ctx context.Context, groupID string, start, end time.Time) ([]Transaction, error)
	FindByType(ctx context.Context, groupID string, tType TransactionType) ([]Transaction, error)
	FindBySource(ctx context.Context, groupID string, souce TransactionSource) ([]Transaction, error)
}

type TransactionService interface {
	Index(ctx context.Context) ([]dto.TransactionResponse, error)
	Show(ctx context.Context, id string) (dto.TransactionResponse, error)
	Create(ctx context.Context, req dto.CreateTransactionRequest, currentUserID string) (dto.TransactionResponse, error)
	Update(ctx context.Context, id string, req dto.UpdateTransactionRequest) (dto.TransactionResponse, error)
	Delete(ctx context.Context, id string) error

	GetByGroup(ctx context.Context, groupID string) ([]dto.TransactionResponse, error)
	GetByDateRange(ctx context.Context, groupID string, start, end time.Time) ([]dto.TransactionResponse, error)
	GetByType(ctx context.Context, groupID string, tType TransactionType) ([]dto.TransactionResponse, error)
	GetBySource(ctx context.Context, groupID string, source TransactionSource) ([]dto.TransactionResponse, error)

	CalculateGroupBalance(ctx context.Context, groupID string) (float64, error)
}
