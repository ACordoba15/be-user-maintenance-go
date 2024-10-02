package models

import (
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model // Nos brinda campos como id, createdAt, updatedAt, deletedAt

	Username string `gorm:"type:varchar(100);not null"`
	Action   string `gorm:"type:varchar(100);not null"`
}
