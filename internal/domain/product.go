package domain

import (
	"errors"
	"time"
)

// Product represents the product entity in the domain layer
type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Price       float64   `json:"price" gorm:"not null"`
	Stock       int       `json:"stock" gorm:"not null;default:0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName specifies the table name for GORM
func (Product) TableName() string {
	return "products"
}

// Validate performs domain-level validation
func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("product name is required")
	}
	if p.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}
	if p.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}
	return nil
}

// IsAvailable checks if the product has sufficient stock
func (p *Product) IsAvailable(quantity int) bool {
	return p.Stock >= quantity
}

// ReduceStock reduces the product stock by the given quantity
func (p *Product) ReduceStock(quantity int) error {
	if !p.IsAvailable(quantity) {
		return errors.New("insufficient stock")
	}
	p.Stock -= quantity
	return nil
}

// BeforeCreate is a GORM hook that runs before creating a product
func (p *Product) BeforeCreate() error {
	return p.Validate()
}
