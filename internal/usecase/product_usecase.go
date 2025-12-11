package usecase

import (
	"context"
	"errors"

	"github.com/example/clean-arch-template/internal/domain"
	"github.com/example/clean-arch-template/internal/repository"
	"gorm.io/gorm"
)

type ProductUseCase struct {
	productRepo repository.ProductRepository
}

func NewProductUseCase(productRepo repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepo: productRepo,
	}
}

// CreateProduct creates a new product
func (uc *ProductUseCase) CreateProduct(ctx context.Context, name, description string, price float64, stock int) (*domain.Product, error) {
	product := &domain.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	if err := uc.productRepo.Create(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// GetProduct retrieves a product by ID
func (uc *ProductUseCase) GetProduct(ctx context.Context, id uint) (*domain.Product, error) {
	product, err := uc.productRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return product, nil
}

// ListProducts retrieves all products
func (uc *ProductUseCase) ListProducts(ctx context.Context) ([]domain.Product, error) {
	return uc.productRepo.FindAll(ctx)
}

// UpdateProduct updates an existing product
func (uc *ProductUseCase) UpdateProduct(ctx context.Context, id uint, name, description string, price float64, stock int) (*domain.Product, error) {
	product, err := uc.productRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	product.Name = name
	product.Description = description
	product.Price = price
	product.Stock = stock

	if err := product.Validate(); err != nil {
		return nil, err
	}

	if err := uc.productRepo.Update(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// DeleteProduct deletes a product
func (uc *ProductUseCase) DeleteProduct(ctx context.Context, id uint) error {
	_, err := uc.productRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	return uc.productRepo.Delete(ctx, id)
}
