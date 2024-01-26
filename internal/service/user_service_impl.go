package service

import (
	"errors"

	"github.com/tasks-hub/users-service/internal/entities"
	"github.com/tasks-hub/users-service/internal/store"
	"golang.org/x/crypto/bcrypt"
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
func (u *UserServiceImpl) CreateUser(input entities.CreateUserInput) (string, error) {
	// Convert CreateUserInput to store.User
	storeUser := &entities.CreateUserInput{
		UserInput: entities.UserInput{
			Username: input.Username,
			Email:    input.Email,
		},
		Password: input.Password,
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(storeUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("can't generate password for user")
	}
	storeUser.Password = string(hashedPassword)

	return u.userStore.CreateUser(storeUser)
}

// GetUserByID retrieves a user by ID
func (u *UserServiceImpl) GetUserByID(userID string) (*entities.User, error) {
	// Call GetUserByID method of UserStore
	user, err := u.userStore.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	user.Password = []byte("")
	return user, nil
}

// UpdateUserProfile updates user profile
func (u *UserServiceImpl) UpdateUserProfile(userID string, input entities.UpdateUserInput) error {
	existingUser, err := u.userStore.GetUserByID(userID)
	if err != nil {
		return err
	}

	existingUser.Username = input.Username
	existingUser.Email = input.Email

	return u.userStore.UpdateUser(existingUser)
}

// ChangePassword changes user password
func (u *UserServiceImpl) ChangePassword(userID string, input entities.UpdateUserPasswordInput) error {
	// Retrieve the existing user for comparison
	existingUser, err := u.userStore.GetUserByID(userID)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(input.OldPassword))
	if err != nil {
		return errors.New("old password is wrong")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("can't generate hash password")
	}

	// Apply changes to the allowed fields
	existingUser.Password = hashedPassword

	// Call the store method to update the user
	return u.userStore.UpdateUser(existingUser)
}

// DeleteUser deletes a user by ID
func (u *UserServiceImpl) DeleteUser(userID string) error {
	user, err := u.userStore.GetUserByID(userID)
	if err != nil {
		return err
	}
	return u.userStore.DeleteUser(user.ID)
}
