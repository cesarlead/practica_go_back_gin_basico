package usecase

import (
	"github.com/cesarlead/practica_go_back_gin_basico/internal/domain"
	"github.com/cesarlead/practica_go_back_gin_basico/internal/repository"
)

// UserUseCase define los casos de uso de Usuarios.
type UserUseCase interface {
	GetAllUsers() ([]*domain.User, error)
	GetUserByID(id int) (*domain.User, error)
	CreateUser(name, email string) (*domain.User, error)
	UpdateUser(id int, name, email string) (*domain.User, error)
	DeleteUser(id int) error
}

type userUseCase struct {
	repo repository.UserRepository
}

// NewUserUseCase inyecta la implementación de repositorio.
func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (uc *userUseCase) GetAllUsers() ([]*domain.User, error) {
	return uc.repo.FindAll()
}

func (uc *userUseCase) GetUserByID(id int) (*domain.User, error) {
	return uc.repo.FindByID(id)
}

func (uc *userUseCase) CreateUser(name, email string) (*domain.User, error) {

	u := &domain.User{Name: name, Email: email}
	id, err := uc.repo.Save(u)
	if err != nil {
		return nil, err
	}
	return uc.repo.FindByID(id)
}

func (uc *userUseCase) UpdateUser(id int, name, email string) (*domain.User, error) {
	u, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	u.Name = name
	u.Email = email
	if err := uc.repo.Update(u); err != nil {
		return nil, err
	}
	return uc.repo.FindByID(id)
}

func (uc *userUseCase) DeleteUser(id int) error {
	return uc.repo.Delete(id)
}
