package usecase

import (
	"context"
	"errors"

	"github.com/example/clean-arch-template/internal/domain"
	"github.com/example/clean-arch-template/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

// Register creates a new user with hashed password
func (uc *UserUseCase) Register(ctx context.Context, email, fullName, password string) (*domain.User, error) {
	// Check if user already exists
	existingUser, _ := uc.userRepo.FindByEmail(ctx, email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:    email,
		FullName: fullName,
		Password: string(hashedPassword),
	}

	// Validate
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// Create user
	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login authenticates a user
func (uc *UserUseCase) Login(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

// GetProfile retrieves user profile by ID
func (uc *UserUseCase) GetProfile(ctx context.Context, userID uint) (*domain.User, error) {
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
