package repository

import (
	"context"

	"github.com/example/clean-arch-template/internal/domain"
)

// ProductRepository defines the interface for product data persistence
type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	FindByID(ctx context.Context, id uint) (*domain.Product, error)
	FindAll(ctx context.Context) ([]domain.Product, error)
	Update(ctx context.Context, product *domain.Product) error
	UpdateStock(ctx context.Context, productID uint, quantity int) error
	Delete(ctx context.Context, id uint) error
}
