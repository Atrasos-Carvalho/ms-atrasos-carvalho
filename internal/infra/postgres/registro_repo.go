package postgres

import (
	"database/sql"
	"ms-atrasos-carvalho/internal/delivery/http/dto"
	"ms-atrasos-carvalho/internal/domain"
	"ms-atrasos-carvalho/internal/domain/repository"
	"time"
)

type RegistroRepo struct {
	db *sql.DB
}

func NewRegistroRepo(db *sql.DB) repository.RegistroRepository {
	return &RegistroRepo{db: db}
}

func (r *RegistroRepo) Create(registro *domain.Registro) error {
	query := "INSERT INTO registro (data, tipo_id) VALUES ($1, $2)"
	_, err := r.db.Exec(query, registro.Data.Time, registro.TipoID)
	return err
}

func (r *RegistroRepo) FindAll() ([]*domain.Registro, error) {
	query := `SELECT id, data, tipo_id FROM registro`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var registros []*domain.Registro

	for rows.Next() {
		reg := &domain.Registro{}
		var rawDate time.Time

		err := rows.Scan(&reg.ID, &rawDate, &reg.TipoID)
		if err != nil {
			return nil, err
		}

		reg.Data = domain.Date{Time: rawDate}
		registros = append(registros, reg)
	}

	return registros, rows.Err()
}

func (r *RegistroRepo) FindByDate(date domain.Date) ([]*domain.Registro, error) {
	query := `SELECT id, data, tipo_id FROM registro WHERE DATE(data) = $1`

	rows, err := r.db.Query(query, date.Time)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var registros []*domain.Registro

	for rows.Next() {
		reg := &domain.Registro{}
		var rawDate time.Time

		err := rows.Scan(&reg.ID, &rawDate, &reg.TipoID)
		if err != nil {
			return nil, err
		}

		reg.Data = domain.Date{Time: rawDate}
		registros = append(registros, reg)
	}

	return registros, rows.Err()
}

func (r *RegistroRepo) FindCount() (dto.ContadorDTO, error) {
	query := `SELECT
		COUNT(CASE WHEN tipo_id = 1 THEN 1 END) AS total_cedo,
		COUNT(CASE WHEN tipo_id = 2 THEN 1 END) AS total_atraso,
		COUNT(CASE WHEN tipo_id = 3 THEN 1 END) AS total_falta
		FROM registro`

	row := r.db.QueryRow(query)

	var contador dto.ContadorDTO
	err := row.Scan(&contador.TotalCedo, &contador.TotalAtraso, &contador.TotalFalta)
	if err != nil {
		return dto.ContadorDTO{}, err
	}

	return contador, nil
}