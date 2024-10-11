package usecase

import (
	"github.com/ACordoba15/be-user-maintenance/internal/domain/models"
	"github.com/ACordoba15/be-user-maintenance/internal/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

// NewUserService crea una nueva instancia de UserService
func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

// GetAllUsers obtiene todos los usuarios
func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

// GetById obtiene un usario por medio del id
func (s *UserService) GetById(id int) (models.User, error) {
	return s.repo.GetById(id)
}

// Login valida si es un usuario válido en la BD
func (s *UserService) Login(username string, password string) (models.User, error) {
	return s.repo.Login(username, password)
}

// AddUser agregar un nuevo usuario
func (s *UserService) AddUser(user models.User) (models.User, error) {
	return s.repo.AddUser(user)
}

// UpdateUser actualiza un usuario
func (s *UserService) UpdateUser(username string, newPassword string) (models.User, error) {
	return s.repo.UpdateUser(username, newPassword)
}

// DeleteUser elimina un usuario de forma lógica
func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
