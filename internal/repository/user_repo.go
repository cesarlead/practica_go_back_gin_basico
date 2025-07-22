package repository

import "github.com/cesarlead/practica_go_back_gin_basico/internal/domain"

// UserRepository abstrae la persistencia de usuarios.
type UserRepository interface {
	FindAll() ([]*domain.User, error)
	FindByID(id int) (*domain.User, error)
	Save(u *domain.User) (int, error) // devuelve el ID insertado
	Update(u *domain.User) error
	Delete(id int) error
}
