package domain

import (
	"errors"
	"time"
)

// User represents the user entity in the domain layer
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	FullName  string    `json:"full_name" gorm:"not null"`
	Password  string    `json:"-" gorm:"not null"` // Never expose password in JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "users"
}

// Validate performs domain-level validation
func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.FullName == "" {
		return errors.New("full name is required")
	}
	if len(u.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}

// BeforeCreate is a GORM hook that runs before creating a user
func (u *User) BeforeCreate() error {
	return u.Validate()
}
