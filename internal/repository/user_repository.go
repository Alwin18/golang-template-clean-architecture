package repository

import (
	"context"

	"github.com/example/clean-arch-template/internal/domain"
)

// UserRepository defines the interface for user data persistence
type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, id uint) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uint) error
}
