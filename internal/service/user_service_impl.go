package service

import (
	"errors"

	"github.com/tasks-hub/users-service/internal/entities"
	"github.com/tasks-hub/users-service/internal/store"
)

// UserServiceImpl contains business logic for users
type UserServiceImpl struct {
	userStore store.UserStore
}

// NewUserService creates a UserServiceImpl instance
func NewUserService(userStore store.UserStore) *UserServiceImpl {
	return &UserServiceImpl{userStore: userStore}
}

// CreateUser registers a new user
func (u *UserServiceImpl) CreateUser(input entities.CreateUserInput) error {
	// Convert CreateUserInput to store.User
	storeUser := &entities.CreateUserInput{
		UserInput: entities.UserInput{
			Username: input.Username,
			Email:    input.Email,
		},
		Password: input.Password,
	}

	return u.userStore.CreateUser(storeUser)
}

// GetUserByID retrieves a user by ID
func (u *UserServiceImpl) GetUserByID(userID int) (*entities.User, error) {
	// Call GetUserByID method of UserStore
	userStore, err := u.userStore.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Convert store.User to entities.User
	return &entities.User{
		ID:       userStore.ID,
		Username: userStore.Username,
		Email:    userStore.Email,
	}, nil
}

// UpdateUserProfile updates user profile
func (u *UserServiceImpl) UpdateUserProfile(userID int, input entities.UpdateUserInput) error {
	existingUser, err := u.userStore.GetUserByID(userID)
	if err != nil {
		return err
	}

	if input.NewPassword != "" {
		return errors.New("update user profile expect a new password")
	}

	existingUser.Username = input.Username
	existingUser.Email = input.Email

	return u.userStore.UpdateUser(existingUser)
}

// ChangePassword changes user password
func (u *UserServiceImpl) ChangePassword(userID int, input entities.UpdateUserInput) error {
	// Retrieve the existing user for comparison
	existingUser, err := u.userStore.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Check for disallowed fields
	if input.Username != "" || input.Email != "" {
		return errors.New("user has not expected fields")
	}

	// Verify old password
	if input.OldPassword != existingUser.Password {
		return errors.New("Old password does not match")
	}

	// Apply changes to the allowed fields
	existingUser.Password = input.NewPassword

	// Call the store method to update the user
	return u.userStore.UpdateUser(existingUser)
}

// DeleteUser deletes a user by ID
func (u *UserServiceImpl) DeleteUser(userID int) error {
	return u.userStore.DeleteUser(userID)
}
