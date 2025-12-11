package handler

import (
	"github.com/example/clean-arch-template/internal/usecase"
	"github.com/example/clean-arch-template/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productUseCase *usecase.ProductUseCase
}

func NewProductHandler(productUseCase *usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"required,gte=0"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"required,gte=0"`
}

// CreateProduct handles product creation
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var req CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	product, err := h.productUseCase.CreateProduct(c.Context(), req.Name, req.Description, req.Price, req.Stock)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Created(c, "Product created successfully", product)
}

// GetProduct retrieves a product by ID
func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid product ID")
	}

	product, err := h.productUseCase.GetProduct(c.Context(), uint(productID))
	if err != nil {
		return response.NotFound(c, err.Error())
	}

	return response.Success(c, "Product retrieved", product)
}

// ListProducts retrieves all products
func (h *ProductHandler) ListProducts(c *fiber.Ctx) error {
	products, err := h.productUseCase.ListProducts(c.Context())
	if err != nil {
		return response.InternalError(c, "Failed to retrieve products")
	}

	return response.Success(c, "Products retrieved", products)
}

// UpdateProduct updates an existing product
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid product ID")
	}

	var req UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	product, err := h.productUseCase.UpdateProduct(c.Context(), uint(productID), req.Name, req.Description, req.Price, req.Stock)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Success(c, "Product updated successfully", product)
}

// DeleteProduct deletes a product
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid product ID")
	}

	if err := h.productUseCase.DeleteProduct(c.Context(), uint(productID)); err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Success(c, "Product deleted successfully", nil)
}
