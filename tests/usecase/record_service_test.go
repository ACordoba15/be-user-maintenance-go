package usecase_test

import (
	"testing"

	"github.com/ACordoba15/be-user-maintenance/internal/domain/models"
	"github.com/ACordoba15/be-user-maintenance/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock de UserRepository para pruebas
type MockRecordRepository struct {
	mock.Mock
}

func (m *MockRecordRepository) GetAll() ([]models.Record, error) {
	args := m.Called()
	return args.Get(0).([]models.Record), args.Error(1)
}

func (m *MockRecordRepository) GetById(id int) (models.Record, error) {
	args := m.Called(id)
	return args.Get(0).(models.Record), args.Error(1)
}

func (m *MockRecordRepository) AddRecord(record models.Record) (models.Record, error) {
	args := m.Called(record)
	return args.Get(0).(models.Record), args.Error(1)
}

func (m *MockRecordRepository) UpdateRecord(action string, id int) (models.Record, error) {
	args := m.Called(action, id)
	return args.Get(0).(models.Record), args.Error(1)
}

func (m *MockRecordRepository) DeleteRecord(id int) error {
	args := m.Called(id)
	return args.Error(1)
}

func TestGetAllRecords(t *testing.T) {
	mockRepo := new(MockRecordRepository)
	recordService := usecase.NewRecordService(mockRepo)

	mockRecords := []models.Record{
		{Username: "userTest1", Action: "TestAction1"},
		{Username: "userTest2", Action: "TestAction2"},
	}

	// Configurar el mock
	mockRepo.On("GetAll").Return(mockRecords, nil)

	records, err := recordService.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, len(records), 2)
	assert.Equal(t, records[0].Username, "userTest1")
	assert.Equal(t, records[0].Action, "TestAction1")

	mockRepo.AssertExpectations(t)
}

func TestGetRecordById(t *testing.T) {
	mockRepo := new(MockRecordRepository)
	recordService := usecase.NewRecordService(mockRepo)

	// Caso de éxito
	t.Run("Success", func(t *testing.T) {
		mockRecord := models.Record{Username: "user1"}

		mockRepo.On("GetById", int(1)).Return(mockRecord, nil)

		record, err := recordService.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, mockRecord.Username, record.Username)

		mockRepo.AssertExpectations(t)
	})
}

func TestAddRecord(t *testing.T) {
	mockRepo := new(MockRecordRepository)
	recordService := usecase.NewRecordService(mockRepo)

	// Caso de éxito
	t.Run("Success", func(t *testing.T) {
		mockRecord := models.Record{Username: "user", Action: "test"}
		mockRecordRequest := models.Record{Username: "user", Action: "test"}

		mockRepo.On("AddRecord", mockRecordRequest).Return(mockRecord, nil)

		record, err := recordService.AddRecord(mockRecordRequest)

		assert.NoError(t, err)
		assert.Equal(t, mockRecord.Username, record.Username)
		assert.Equal(t, mockRecord.Action, record.Action)

		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateRecord(t *testing.T) {
	mockRepo := new(MockRecordRepository)
	recordService := usecase.NewRecordService(mockRepo)

	// Caso de éxito
	t.Run("Success", func(t *testing.T) {
		mockRecord := models.Record{Username: "user", Action: "test"}

		mockRepo.On("UpdateRecord", "test", 1).Return(mockRecord, nil)

		record, err := recordService.UpdateRecord("test", 1)

		assert.NoError(t, err)
		assert.Equal(t, mockRecord.Action, record.Action)

		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteRecord(t *testing.T) {
	mockRepo := new(MockRecordRepository)
	recordService := usecase.NewRecordService(mockRepo)

	// Caso de éxito
	t.Run("Success", func(t *testing.T) {
		mockRepo.On("DeleteRecord", 1).Return(nil, nil)

		err := recordService.DeleteRecord(1)

		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})
}
