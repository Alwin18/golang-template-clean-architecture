package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/example/clean-arch-template/internal/domain"
	"github.com/example/clean-arch-template/internal/infrastructure/persistence"
	"gorm.io/gorm"
)

// CreateOrderRequest represents the request to create an order
type CreateOrderRequest struct {
	UserID        uint                     `json:"user_id"`
	PaymentMethod string                   `json:"payment_method"`
	Items         []CreateOrderItemRequest `json:"items"`
}

// CreateOrderItemRequest represents an item in the order
type CreateOrderItemRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type OrderUseCase struct {
	db *gorm.DB
}

func NewOrderUseCase(db *gorm.DB) *OrderUseCase {
	return &OrderUseCase{
		db: db,
	}
}

// CreateOrder creates a new order with transaction support
// This is the KEY EXAMPLE of multi-table transaction with GORM
func (uc *OrderUseCase) CreateOrder(ctx context.Context, req CreateOrderRequest) (*domain.Order, error) {
	var createdOrder *domain.Order

	// Start GORM Transaction
	err := uc.db.Transaction(func(tx *gorm.DB) error {
		// Create repository instances with transaction
		orderRepo := persistence.NewOrderRepository(tx)
		productRepo := persistence.NewProductRepository(tx)
		paymentRepo := persistence.NewPaymentRepository(tx)

		// Step 1: Validate and prepare order items
		var orderItems []domain.OrderItem
		var totalAmount float64

		for _, item := range req.Items {
			// Fetch product within transaction
			product, err := productRepo.FindByID(ctx, item.ProductID)
			if err != nil {
				return fmt.Errorf("product with ID %d not found", item.ProductID)
			}

			// Check stock availability
			if !product.IsAvailable(item.Quantity) {
				return fmt.Errorf("insufficient stock for product %s (available: %d, requested: %d)",
					product.Name, product.Stock, item.Quantity)
			}

			// Calculate price
			itemTotal := product.Price * float64(item.Quantity)
			totalAmount += itemTotal

			// Prepare order item
			orderItems = append(orderItems, domain.OrderItem{
				ProductID: product.ID,
				Quantity:  item.Quantity,
				Price:     product.Price,
			})

			// Step 2: Reduce product stock
			if err := product.ReduceStock(item.Quantity); err != nil {
				return err
			}
			if err := productRepo.Update(ctx, product); err != nil {
				return fmt.Errorf("failed to update product stock: %w", err)
			}
		}

		// Step 3: Create order
		order := &domain.Order{
			UserID:      req.UserID,
			TotalAmount: totalAmount,
			Status:      domain.OrderStatusPending,
			Items:       orderItems,
		}

		if err := order.Validate(); err != nil {
			return fmt.Errorf("order validation failed: %w", err)
		}

		if err := orderRepo.Create(ctx, order); err != nil {
			return fmt.Errorf("failed to create order: %w", err)
		}

		// Step 4: Create payment record
		payment := &domain.Payment{
			OrderID: order.ID,
			Amount:  totalAmount,
			Status:  domain.PaymentStatusPending,
			Method:  req.PaymentMethod,
		}

		if err := payment.Validate(); err != nil {
			return fmt.Errorf("payment validation failed: %w", err)
		}

		if err := paymentRepo.Create(ctx, payment); err != nil {
			return fmt.Errorf("failed to create payment: %w", err)
		}

		createdOrder = order
		return nil // Commit transaction if all operations succeed
	})

	if err != nil {
		return nil, err // Auto rollback on error
	}

	return createdOrder, nil
}

// GetOrderDetail retrieves order details by ID
func (uc *OrderUseCase) GetOrderDetail(ctx context.Context, orderID uint) (*domain.Order, error) {
	orderRepo := persistence.NewOrderRepository(uc.db)

	order, err := orderRepo.FindByID(ctx, orderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	return order, nil
}

// ListUserOrders retrieves all orders for a specific user
func (uc *OrderUseCase) ListUserOrders(ctx context.Context, userID uint) ([]domain.Order, error) {
	orderRepo := persistence.NewOrderRepository(uc.db)
	return orderRepo.FindByUserID(ctx, userID)
}
