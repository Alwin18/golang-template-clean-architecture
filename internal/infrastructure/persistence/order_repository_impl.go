package persistence

import (
	"context"

	"github.com/example/clean-arch-template/internal/domain"
	"github.com/example/clean-arch-template/internal/repository"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new instance of OrderRepository
// Accepts *gorm.DB which can be either a regular connection or a transaction
func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(ctx context.Context, order *domain.Order) error {
	// GORM will automatically create order items due to association
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *orderRepository) FindByID(ctx context.Context, id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.WithContext(ctx).
		Preload("Items").
		Preload("Items.Product").
		Preload("User").
		First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) FindByUserID(ctx context.Context, userID uint) ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.WithContext(ctx).
		Preload("Items").
		Preload("Items.Product").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).Error
	return orders, err
}

func (r *orderRepository) Update(ctx context.Context, order *domain.Order) error {
	return r.db.WithContext(ctx).Save(order).Error
}
