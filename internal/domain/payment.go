package domain

import (
	"errors"
	"time"
)

// PaymentStatus represents the status of a payment
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
)

// Payment represents the payment entity in the domain layer
type Payment struct {
	ID        uint          `json:"id" gorm:"primaryKey"`
	OrderID   uint          `json:"order_id" gorm:"not null;uniqueIndex"`
	Order     *Order        `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Amount    float64       `json:"amount" gorm:"not null"`
	Status    PaymentStatus `json:"status" gorm:"not null;default:'pending'"`
	Method    string        `json:"method" gorm:"not null"` // e.g., "credit_card", "bank_transfer"
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// TableName specifies the table name for GORM
func (Payment) TableName() string {
	return "payments"
}

// Validate performs domain-level validation
func (p *Payment) Validate() error {
	if p.OrderID == 0 {
		return errors.New("order ID is required")
	}
	if p.Amount <= 0 {
		return errors.New("payment amount must be greater than 0")
	}
	if p.Method == "" {
		return errors.New("payment method is required")
	}
	return nil
}

// MarkAsCompleted marks the payment as completed
func (p *Payment) MarkAsCompleted() {
	p.Status = PaymentStatusCompleted
}

// MarkAsFailed marks the payment as failed
func (p *Payment) MarkAsFailed() {
	p.Status = PaymentStatusFailed
}

// BeforeCreate is a GORM hook that runs before creating a payment
func (p *Payment) BeforeCreate() error {
	return p.Validate()
}
