package domain

import (
	"context"
	"database/sql"
	"time"
)

type TransactionType string
type TransactionSource string

const (
	Income TransactionType = "income"
	Expense TransactionType = "expense"

	Manual TransactionSource = "manual"
	GmailBrimo TransactionSource = "gmail_brimo"
	ImportCSV TransactionSource = "import_csv"
	QRScan TransactionSource = "qr_scan"
)

type Transaction struct {
	ID          string  `db:"id"`
	GroupID     string  `db:"group_id"`
	CreatedBy  string  `db:"created_by"`
	PaidBy     string  `db:"paid_by"`
	Amount      float64 `db:"amount"`
	Description sql.NullString `db:"description"`
	Date time.Time `db:"date"`
	Type TransactionType `db:"type"`
	Source TransactionSource `db:"source"`
	ExternalRef sql.NullString `db:"external_ref"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type TransactionRepository interface {
	FindByID(ctx context.Context, id string) (Transaction, error)
	FindByGroupID(ctx context.Context, groupID string) ([]Transaction, error)
	GetAll(ctx context.Context) ([]Transaction, error)
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
	Index(ctx context.Context) ([]Transaction, error)
	Create(ctx context.Context, req Transaction) (Transaction, error)
	Update(ctx context.Context, id string, req Transaction) (Transaction, error)
	Delete(ctx context.Context, id string) error
	Show(ctx context.Context, id string) (Transaction, error)

	GetByGroup(ctx context.Context, groupID string) ([]Transaction, error)
	GetByDateRange(ctx context.Context, groupID string, start, end time.Time) ([]Transaction, error)
	GetByType(ctx context.Context, groupID string, tType TransactionType) ([]Transaction, error)
	GetBySource(ctx context.Context, groupID string, source TransactionSource) ([]Transaction, error)

	CalculateGroupBalance(ctx context.Context, groupID string) (float64, error)
}