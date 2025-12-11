package repository

import (
	"context"

	"github.com/example/clean-arch-template/internal/domain"
)

// PaymentRepository defines the interface for payment data persistence
type PaymentRepository interface {
	Create(ctx context.Context, payment *domain.Payment) error
	FindByID(ctx context.Context, id uint) (*domain.Payment, error)
	FindByOrderID(ctx context.Context, orderID uint) (*domain.Payment, error)
	Update(ctx context.Context, payment *domain.Payment) error
}
