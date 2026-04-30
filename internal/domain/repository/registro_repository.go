package repository

import "ms-atrasos-carvalho/internal/domain"

type RegistroRepository interface {
	Create(registro *domain.Registro) error
	FindAll() ([]*domain.Registro, error)
	FindByDate(date domain.Date) ([]*domain.Registro, error)
}