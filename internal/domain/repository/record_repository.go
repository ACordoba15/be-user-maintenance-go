package repository

import (
	"github.com/ACordoba15/be-user-maintenance/db"
	"github.com/ACordoba15/be-user-maintenance/internal/domain/models"
	"gorm.io/gorm"
)

// RecordRepository define la interfaz para el repositorio de records
type RecordRepository interface {
	GetAll() ([]models.Record, error)
	GetById(id int) (models.Record, error)
	AddRecord(record models.Record) (models.Record, error)
	UpdateRecord(action string, id int) (models.Record, error)
	DeleteRecord(id int) error
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

// Obtiene un registro por id
func (r *recordRepository) GetById(id int) (models.Record, error) {
	var record models.Record
	result := r.db.First(&record, id).Error

	if result != nil {
		return models.Record{}, result
	}

	return record, nil
}

// Agrega un nuevo usuario
func (r *recordRepository) AddRecord(user models.Record) (models.Record, error) {
	result := r.db.Create(&user).Error

	if result != nil {
		return models.Record{}, result
	}

	return user, result
}

// Actualiza un usuario existente por usename
func (r *recordRepository) UpdateRecord(action string, id int) (models.Record, error) {
	var record models.Record
	result := r.db.First(&record, id).Error

	if result != nil {
		return models.Record{}, result
	}

	record.Action = action
	result = db.DB.Save(&record).Error
	return record, result
}

// Elimina un registro por id
func (r *recordRepository) DeleteRecord(id int) error {
	result := r.db.Delete(&models.Record{}, id).Error
	return result
}
