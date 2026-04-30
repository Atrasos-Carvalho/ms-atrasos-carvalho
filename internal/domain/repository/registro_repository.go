package repository

import (
	"ms-atrasos-carvalho/internal/delivery/http/dto"
	"ms-atrasos-carvalho/internal/domain"
)

type RegistroRepository interface {
	Create(registro *domain.Registro) error
	FindAll() ([]*domain.Registro, error)
	FindByDate(date domain.Date) ([]*domain.Registro, error)
	FindCount() (dto.ContadorDTO, error)
}