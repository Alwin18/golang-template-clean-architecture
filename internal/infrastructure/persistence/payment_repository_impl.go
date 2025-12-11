package persistence

import (
	"context"

	"github.com/example/clean-arch-template/internal/domain"
	"github.com/example/clean-arch-template/internal/repository"
	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

// NewPaymentRepository creates a new instance of PaymentRepository
// Accepts *gorm.DB which can be either a regular connection or a transaction
func NewPaymentRepository(db *gorm.DB) repository.PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(ctx context.Context, payment *domain.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

func (r *paymentRepository) FindByID(ctx context.Context, id uint) (*domain.Payment, error) {
	var payment domain.Payment
	err := r.db.WithContext(ctx).
		Preload("Order").
		First(&payment, id).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) FindByOrderID(ctx context.Context, orderID uint) (*domain.Payment, error) {
	var payment domain.Payment
	err := r.db.WithContext(ctx).
		Preload("Order").
		Where("order_id = ?", orderID).
		First(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) Update(ctx context.Context, payment *domain.Payment) error {
	return r.db.WithContext(ctx).Save(payment).Error
}
