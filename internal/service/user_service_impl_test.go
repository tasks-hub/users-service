package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tasks-hub/users-service/internal/entities"
	"github.com/tasks-hub/users-service/internal/service"
)

// Mock UserStore for testing
type MockUserStore struct {
	mock.Mock
}

func (m *MockUserStore) CreateUser(input *entities.CreateUserInput) (*entities.User, error) {
	args := m.Called(input)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserStore) GetUserByID(userID string) (*entities.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserStore) GetUserByEmail(userCredentials *entities.UserCredentials) (*entities.User, error) {
	args := m.Called(userCredentials)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserStore) UpdateUser(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserStore) DeleteUser(userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}

// MockPasswordHasher is a mock implementation of the PasswordHasher interface
type MockPasswordHasher struct {
	mock.Mock
}

// GenerateHash mocks the GenerateHash method of the PasswordHasher interface
func (m *MockPasswordHasher) GenerateHash(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

// GenerateHash mocks the CompareHashAndPassword method of the PasswordHasher interface
func (m *MockPasswordHasher) CompareHashAndPassword(hash, password []byte) error {
	args := m.Called(password)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	mockUserStore := new(MockUserStore)
	mockPasswordHasher := new(MockPasswordHasher)
	userService := service.NewUserService(mockUserStore, mockPasswordHasher)

	testInput := entities.CreateUserInput{
		UserInput: entities.UserInput{
			Username: "testuser",
			Email:    "test@example.com",
		},
		Password: "testpassword",
	}

	expectedUser := &entities.User{
		ID:       "some-id",
		Username: testInput.Username,
		Email:    testInput.Email,
		Password: []byte(testInput.Password),
	}

	// Mock the password hashing to return the same password
	mockPasswordHasher.On("GenerateHash", testInput.Password).Return(string(expectedUser.Password), nil)
	// Ensure that CreateUser is called with the exact testInput (including the hashed password)
	mockUserStore.On("CreateUser", &testInput).Return(expectedUser, nil)

	user, err := userService.CreateUser(testInput)

	mockUserStore.AssertExpectations(t)
	mockPasswordHasher.AssertExpectations(t)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.ID, user.ID)
	// Add more assertions based on your business logic and expected behavior
}

func TestGetUserByID(t *testing.T) {
	mockUserStore := new(MockUserStore)
	mockPasswordHasher := new(MockPasswordHasher)
	userService := service.NewUserService(mockUserStore, mockPasswordHasher)

	testUserID := "some-id"
	expectedUser := &entities.User{
		ID:       testUserID,
		Username: "testuser",
		Email:    "test@example.com",
		Password: []byte("hashedpassword"),
	}

	mockUserStore.On("GetUserByID", testUserID).Return(expectedUser, nil)

	user, err := userService.GetUserByID(testUserID)

	mockUserStore.AssertExpectations(t)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.ID, user.ID)
	// Add more assertions based on your business logic and expected behavior
}
