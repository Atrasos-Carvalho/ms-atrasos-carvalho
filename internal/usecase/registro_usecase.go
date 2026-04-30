package usecase

import (
	"ms-atrasos-carvalho/internal/domain"
	"ms-atrasos-carvalho/internal/domain/repository"
)

type RegistroUseCase struct {
	repo repository.RegistroRepository
}

func NewRegistroUseCase(repo repository.RegistroRepository) *RegistroUseCase {
	return &RegistroUseCase{repo: repo}
}

func (uc *RegistroUseCase) Create(registro *domain.Registro) error {
	return uc.repo.Create(registro)
}

func (uc *RegistroUseCase) FindAll() ([]*domain.Registro, error) {
	return uc.repo.FindAll()
}

func (uc *RegistroUseCase) FindByDate(date domain.Date) ([]*domain.Registro, error) {
	return uc.repo.FindByDate(date)
}