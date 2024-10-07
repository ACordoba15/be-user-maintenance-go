package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `swaggerignore:"true"` // Nos brinda campos como id, createdAt, updatedAt, deletedAt

	Username string `gorm:"type:varchar(100);not null;unique_index"`
	Password string `gorm:"varchar(200);not null;"`
}
