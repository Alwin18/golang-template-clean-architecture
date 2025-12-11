package handler

import (
	"github.com/example/clean-arch-template/internal/usecase"
	"github.com/example/clean-arch-template/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	orderUseCase *usecase.OrderUseCase
}

func NewOrderHandler(orderUseCase *usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: orderUseCase,
	}
}

// CreateOrder handles order creation with transaction
func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var req usecase.CreateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	// Create order (with automatic transaction handling)
	order, err := h.orderUseCase.CreateOrder(c.Context(), req)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Created(c, "Order created successfully", order)
}

// GetOrderDetail retrieves order details
func (h *OrderHandler) GetOrderDetail(c *fiber.Ctx) error {
	orderID, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid order ID")
	}

	order, err := h.orderUseCase.GetOrderDetail(c.Context(), uint(orderID))
	if err != nil {
		return response.NotFound(c, err.Error())
	}

	return response.Success(c, "Order retrieved", order)
}

// ListUserOrders retrieves all orders for a user
func (h *OrderHandler) ListUserOrders(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id")
	if err != nil {
		return response.BadRequest(c, "Invalid user ID")
	}

	orders, err := h.orderUseCase.ListUserOrders(c.Context(), uint(userID))
	if err != nil {
		return response.InternalError(c, "Failed to retrieve orders")
	}

	return response.Success(c, "Orders retrieved", orders)
}
