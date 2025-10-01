package service

import (
	"context"
	"divvy/divvy-api/domain"
	"divvy/divvy-api/dto"
	"errors"
	"time"

	"github.com/google/uuid"
)

type WalletService struct {
	WalletRepository domain.WalletRepository
}

func NewWallet(walletRepository domain.WalletRepository) domain.WalletService {
	return &WalletService{
		WalletRepository: walletRepository,
	}
}

// CreateWallet implements domain.WalletService.
func (w *WalletService) CreateWallet(ctx context.Context, userID string, req dto.CreateWalletRequest) (dto.WalletResponse, error) {
	if req.Balance < 0 {
		return dto.WalletResponse{}, errors.New("initial balance cannot be negative")
	}

	wallet := domain.Wallet{
		ID:        uuid.NewString(),
		User_id:   userID,
		Balance:   req.Balance,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := w.WalletRepository.Save(ctx, &wallet); err != nil {
		return dto.WalletResponse{}, err
	}

	return dto.WalletResponse{
		ID:        wallet.ID,
		UserID:    wallet.User_id,
		Balance:   wallet.Balance,
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}, nil
}

// DeleteWallet implements domain.WalletService.
func (w *WalletService) DeleteWallet(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetWalletByID implements domain.WalletService.
func (w *WalletService) GetWalletByID(ctx context.Context, id string) (dto.WalletResponse, error) {
	wallet, err := w.WalletRepository.FindById(ctx, id)
	if err != nil {
		return dto.WalletResponse{}, err
	}
	if wallet.ID == "" {
		return dto.WalletResponse{}, errors.New("wallet not found")
	}
	return dto.WalletResponse{
		ID:        wallet.ID,
		UserID:    wallet.User_id,
		Balance:   wallet.Balance,
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}, nil
}

// GetWalletByUser implements domain.WalletService.
func (w *WalletService) GetWalletByUser(ctx context.Context, uid string) ([]dto.WalletResponse, error) {
	panic("unimplemented")
}

// GetWalletWithTransactions implements domain.WalletService.
func (w *WalletService) GetWalletWithTransactions(ctx context.Context, walletID string) (dto.WalletWithTransactionsResponse, error) {
	panic("unimplemented")
}

// GetWalletsWithTransactions implements domain.WalletService.
func (w *WalletService) GetWalletsWithTransactions(ctx context.Context, userID string) ([]dto.WalletWithTransactionsResponse, error) {
	panic("unimplemented")
}

// UpdateWallet implements domain.WalletService.
func (w *WalletService) UpdateWallet(ctx context.Context, id string, req dto.UpdateWalletRequest) (dto.WalletResponse, error) {
	existing, err := w.WalletRepository.FindById(ctx, id)
	if err != nil {
		return dto.WalletResponse{}, err
	}

	if existing.ID == "" {
		return dto.WalletResponse{}, errors.New("wallet not found")
	}

	if req.Balance != 0 {
		existing.Balance = req.Balance
	}

	existing.UpdatedAt = time.Now()

	wallet, err := w.WalletRepository.Update(ctx, &existing)
	if err != nil {
		return dto.WalletResponse{}, err
	}

	return dto.WalletResponse{
		ID: wallet.ID,
		UserID: wallet.User_id,
		Balance: wallet.Balance,
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}, nil
}
