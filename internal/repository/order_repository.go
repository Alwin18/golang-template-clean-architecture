package repository

import (
	"context"

	"github.com/example/clean-arch-template/internal/domain"
)

// OrderRepository defines the interface for order data persistence
type OrderRepository interface {
	Create(ctx context.Context, order *domain.Order) error
	FindByID(ctx context.Context, id uint) (*domain.Order, error)
	FindByUserID(ctx context.Context, userID uint) ([]domain.Order, error)
	Update(ctx context.Context, order *domain.Order) error
}
