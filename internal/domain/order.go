package domain

import (
	"errors"
	"time"
)

// OrderStatus represents the status of an order
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusCancelled OrderStatus = "cancelled"
	OrderStatusCompleted OrderStatus = "completed"
)

// Order represents the order entity in the domain layer
type Order struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	UserID      uint        `json:"user_id" gorm:"not null;index"`
	User        *User       `json:"user,omitempty" gorm:"foreignKey:UserID"`
	TotalAmount float64     `json:"total_amount" gorm:"not null"`
	Status      OrderStatus `json:"status" gorm:"not null;default:'pending'"`
	Items       []OrderItem `json:"items" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// TableName specifies the table name for GORM
func (Order) TableName() string {
	return "orders"
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderID   uint      `json:"order_id" gorm:"not null;index"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Product   *Product  `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null"` // Price at the time of order
	CreatedAt time.Time `json:"created_at"`
}

// TableName specifies the table name for GORM
func (OrderItem) TableName() string {
	return "order_items"
}

// Validate performs domain-level validation for Order
func (o *Order) Validate() error {
	if o.UserID == 0 {
		return errors.New("user ID is required")
	}
	if len(o.Items) == 0 {
		return errors.New("order must have at least one item")
	}
	if o.TotalAmount <= 0 {
		return errors.New("total amount must be greater than 0")
	}
	return nil
}

// CalculateTotal calculates the total amount based on items
func (o *Order) CalculateTotal() {
	var total float64
	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}
	o.TotalAmount = total
}

// ValidateItem validates an order item
func (item *OrderItem) Validate() error {
	if item.ProductID == 0 {
		return errors.New("product ID is required")
	}
	if item.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	if item.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	return nil
}
