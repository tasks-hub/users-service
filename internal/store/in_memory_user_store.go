// internal/store/user_store.go

package store

import (
	"errors"
	"sync"

	"github.com/tasks-hub/users-service/internal/entities"
)

// InMemoryUserStore is a fictional implementation of UserStore using an in-memory database
type InMemoryUserStore struct {
	users map[int]*entities.User
	mu    sync.RWMutex
}

// NewInMemoryUserStore creates a new instance of InMemoryUserStore
func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[int]*entities.User),
	}
}

// GetUserByID retrieves a user by ID from the in-memory database
func (u *InMemoryUserStore) GetUserByID(userID int) (*entities.User, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	user, exists := u.users[userID]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// CreateUser creates a new user in the in-memory database
func (u *InMemoryUserStore) CreateUser(user *entities.CreateUserInput) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	// Assign a new ID (in a real implementation this should be a task for database)
	newUserID := len(u.users) + 1

	// Create a new user entity
	newUser := &entities.User{
		ID:       newUserID,
		Username: user.Username,
		Email:    user.Email,
		// Other fields can be initialized here
	}

	u.users[newUserID] = newUser

	return nil
}

// UpdateUser updates the information of an existing user in the in-memory database
func (u *InMemoryUserStore) UpdateUser(updatedUser *entities.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	// Check if the user exists
	user, exists := u.users[updatedUser.ID]
	if !exists {
		return errors.New("user not found")
	}

	// Update only allowed fields
	user.Username = updatedUser.Username
	user.Email = updatedUser.Email
	user.Password = updatedUser.Password

	return nil
}

// DeleteUser deletes a user by ID from the in-memory database
func (u *InMemoryUserStore) DeleteUser(userID int) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	// Check if the user exists
	_, exists := u.users[userID]
	if !exists {
		return errors.New("user not found")
	}

	// Delete the user from the in-memory database
	delete(u.users, userID)

	return nil
}
