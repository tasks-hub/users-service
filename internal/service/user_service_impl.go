package service

import (
	"errors"
	"fmt"

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
func (u *UserServiceImpl) CreateUser(input entities.CreateUserInput) (*entities.User, error) {
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
		return nil, errors.New("can't generate password for user")
	}
	storeUser.Password = string(hashedPassword)

	user, err := u.userStore.CreateUser(storeUser)
	if err != nil {
		return nil, err
	}
	user.Password = []byte("")
	return user, nil
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

// GetUserByEmail retrieves a user by email and password
func (u *UserServiceImpl) GetUserByEmail(userCredentials *entities.UserCredentials) (*entities.User, error) {
	// Call GetUserByEmail method of UserStore
	user, err := u.userStore.GetUserByEmail(userCredentials)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(userCredentials.Password)); err != nil {
		return nil, errors.New(fmt.Sprintf("wrong password for user %s", user.Email))
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
