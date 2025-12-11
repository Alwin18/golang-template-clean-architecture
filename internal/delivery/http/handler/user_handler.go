package handler

import (
	"github.com/example/clean-arch-template/internal/usecase"
	"github.com/example/clean-arch-template/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Register handles user registration
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	user, err := h.userUseCase.Register(c.Context(), req.Email, req.FullName, req.Password)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Created(c, "User registered successfully", user)
}

// Login handles user authentication
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	user, err := h.userUseCase.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return response.Unauthorized(c, err.Error())
	}

	return response.Success(c, "Login successful", user)
}

// GetProfile retrieves user profile
func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid user ID")
	}

	user, err := h.userUseCase.GetProfile(c.Context(), uint(userID))
	if err != nil {
		return response.NotFound(c, err.Error())
	}

	return response.Success(c, "User profile retrieved", user)
}
