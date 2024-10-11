package usecase_test

import (
	"testing"

	"github.com/ACordoba15/be-user-maintenance/internal/domain/models"
	"github.com/ACordoba15/be-user-maintenance/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock de UserRepository para pruebas
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAll() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) GetById(id int) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) Login(username string, password string) (models.User, error) {
	args := m.Called(username, password)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) AddUser(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(username string, newPassword string) (models.User, error) {
	args := m.Called(username, newPassword)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(1)
}

func TestGetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := usecase.NewUserService(mockRepo)

	mockUsers := []models.User{
		{Username: "userTest1", Password: "passwordTest1"},
		{Username: "userTest2", Password: "passwordTest2"},
	}

	// Configurar el mock
	mockRepo.On("GetAll").Return(mockUsers, nil)

	users, err := userService.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, len(users), 2)
	assert.Equal(t, users[0].Username, "userTest1")

	mockRepo.AssertExpectations(t)
}

func TestGetById(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := usecase.NewUserService(mockRepo)

	// Caso de éxito
	t.Run("Success", func(t *testing.T) {
		mockUser := models.User{Username: "user1"}

		mockRepo.On("GetById", int(1)).Return(mockUser, nil)

		user, err := userService.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, mockUser.Username, user.Username)

		mockRepo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := usecase.NewUserService(mockRepo)

	// Caso de éxito
	t.Run("Success", func(t *testing.T) {
		mockUser := models.User{Username: "user", Password: "test"}

		mockRepo.On("Login", string("user"), string("test")).Return(mockUser, nil)

		user, err := userService.Login("user", "test")

		assert.NoError(t, err)
		assert.Equal(t, mockUser.Username, user.Username)
		assert.Equal(t, mockUser.Password, user.Password)

		mockRepo.AssertExpectations(t)
	})
}

func TestAddUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := usecase.NewUserService(mockRepo)

	// Caso de éxito
	t.Run("Success", func(t *testing.T) {
		mockUser := models.User{Username: "user", Password: "test"}
		mockUserRequest := models.User{Username: "user", Password: "test"}

		mockRepo.On("AddUser", mockUserRequest).Return(mockUser, nil)

		user, err := userService.AddUser(mockUserRequest)

		assert.NoError(t, err)
		assert.Equal(t, mockUser.Username, user.Username)
		assert.Equal(t, mockUser.Password, user.Password)

		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := usecase.NewUserService(mockRepo)

	// Caso de éxito
	t.Run("Success", func(t *testing.T) {
		mockUser := models.User{Username: "user", Password: "newPass"}

		mockRepo.On("UpdateUser", "user", "newPass").Return(mockUser, nil)

		user, err := userService.UpdateUser("user", "newPass")

		assert.NoError(t, err)
		assert.Equal(t, mockUser.Username, user.Username)
		assert.Equal(t, mockUser.Password, user.Password)

		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := usecase.NewUserService(mockRepo)

	// Caso de éxito
	t.Run("Success", func(t *testing.T) {
		mockRepo.On("DeleteUser", 1).Return(nil, nil)

		err := userService.DeleteUser(1)

		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})
}
