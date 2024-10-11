package usecase

import (
	"github.com/ACordoba15/be-user-maintenance/internal/domain/models"
	"github.com/ACordoba15/be-user-maintenance/internal/domain/repository"
)

type RecordService struct {
	repo repository.RecordRepository
}

func NewRecordService(r repository.RecordRepository) *RecordService {
	return &RecordService{repo: r}
}

// Obtiene todos los registros
func (s *RecordService) GetAll() ([]models.Record, error) {
	return s.repo.GetAll()
}

// GetById obtiene un registro por medio del id
func (s *RecordService) GetById(id int) (models.Record, error) {
	return s.repo.GetById(id)
}

// AddRecord agrega un nuevo registro
func (s *RecordService) AddRecord(record models.Record) (models.Record, error) {
	return s.repo.AddRecord(record)
}

// UpdateRecord actualiza un registro
func (s *RecordService) UpdateRecord(action string, id int) (models.Record, error) {
	return s.repo.UpdateRecord(action, id)
}

// DeleteRecord elimina un registro de forma lógica
func (s *RecordService) DeleteRecord(id int) error {
	return s.repo.DeleteRecord(id)
}
