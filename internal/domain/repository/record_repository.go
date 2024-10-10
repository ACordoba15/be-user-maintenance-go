package repository

import (
	"github.com/ACordoba15/be-user-maintenance/internal/domain/models"
	"gorm.io/gorm"
)

// RecordRepository define la interfaz para el repositorio de records
type RecordRepository interface {
	GetAll() ([]models.Record, error)
	// Otros métodos como Save, FindByID, etc.
}

// Recordtory es la implementación del RecordRepository
type recordRepository struct {
	db *gorm.DB
}

// NewRecordRepository crea una nueva instancia de RecordRepository
func NewRecordRepository(db *gorm.DB) RecordRepository {
	return &recordRepository{db: db}
}

// GetAll obtiene todos los registros de la bitácora de la base de datos
func (r *recordRepository) GetAll() ([]models.Record, error) {
	var records []models.Record
	if err := r.db.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}
