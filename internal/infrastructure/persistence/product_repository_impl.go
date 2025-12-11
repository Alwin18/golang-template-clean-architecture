package persistence

import (
	"context"

	"github.com/example/clean-arch-template/internal/domain"
	"github.com/example/clean-arch-template/internal/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of ProductRepository
// Accepts *gorm.DB which can be either a regular connection or a transaction
func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, product *domain.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *productRepository) FindByID(ctx context.Context, id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.WithContext(ctx).First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FindAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.WithContext(ctx).Find(&products).Error
	return products, err
}

func (r *productRepository) Update(ctx context.Context, product *domain.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

// UpdateStock updates product stock with pessimistic locking to prevent race conditions
func (r *productRepository) UpdateStock(ctx context.Context, productID uint, quantity int) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var product domain.Product

		// Use FOR UPDATE lock to prevent concurrent updates
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&product, productID).Error; err != nil {
			return err
		}

		// Reduce stock
		if err := product.ReduceStock(quantity); err != nil {
			return err
		}

		// Save updated stock
		return tx.Save(&product).Error
	})
}

func (r *productRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Product{}, id).Error
}
