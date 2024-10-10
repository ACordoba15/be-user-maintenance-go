package repository

import (
	"github.com/ACordoba15/be-user-maintenance/db"
	"github.com/ACordoba15/be-user-maintenance/internal/domain/models"
	"gorm.io/gorm"
)

// UserRepository define la interfaz para el repositorio de usuarios
type UserRepository interface {
	GetAll() ([]models.User, error)
	GetById(id int) (models.User, error)
	Login(username string, password string) (models.User, error)
	AddUser(user models.User) (models.User, error)
	UpdateUser(username string, password string) (models.User, error)
	DeleteUser(id int) error
	// Otros métodos como Save, FindByID, etc.
}

// userRepository es la implementación del UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository crea una nueva instancia de userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// GetAll obtiene todos los usuarios de la base de datos
func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users).Error

	if result != nil {
		return []models.User{}, result
	}

	return users, nil
}

// Obtiene un usurio por id
func (r *userRepository) GetById(id int) (models.User, error) {
	var user models.User
	result := r.db.First(&user, id).Error

	if result != nil {
		return models.User{}, result
	}

	return user, nil
}

// Valida un usuario en la BD
func (r *userRepository) Login(username string, password string) (models.User, error) {
	var user models.User
	result := r.db.Where("username = ? and password = ?", username, password).First(&user).Error

	if result != nil {
		return models.User{}, result
	}

	return user, result
}

// Agrega un nuevo usuario
func (r *userRepository) AddUser(user models.User) (models.User, error) {
	result := r.db.Create(&user).Error

	if result != nil {
		return models.User{}, result
	}

	return user, result
}

// Actualiza un usuario existente por usename
func (r *userRepository) UpdateUser(username string, newPassword string) (models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return models.User{}, err
	}

	user.Password = newPassword
	result := db.DB.Save(&user).Error
	return user, result
}

// Elimina un usuario por id
func (r *userRepository) DeleteUser(id int) error {
	result := r.db.Delete(&models.User{}, id).Error
	return result
}
