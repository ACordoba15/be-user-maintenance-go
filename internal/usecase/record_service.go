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

func (s *RecordService) GetAllRecords() ([]models.Record, error) {
	return s.repo.GetAll()
}
