package entities

import (
	"time"
)

// UserInput represents the common input for user-related operations
type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// CreateUserInput represents the input for user registration
type CreateUserInput struct {
	UserInput
	Password string `json:"password" binding:"required,min=12"`
}

// UpdateUserInput represents the input for updating user profile
type UpdateUserInput struct {
	UserID      string `json:"userID"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	OldPassword string `json:"oldPassword" binding:"required,min=8"`
	NewPassword string `json:"newPassword" binding:"required,min=8"`
}

// User represents the user entity
type User struct {
	ID       string  `db:"id"`
	Username string  `db:"username"`
	Email    string  `db:"email"`
	Password []byte  `db:"password"`
	Role     *string `db:"role"`

	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
