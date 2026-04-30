package dto

type ContadorDTO struct {
	TotalCedo int `json:"total_cedo"`
	TotalAtraso int `json:"total_atraso"`
	TotalFalta int `json:"total_falta"`
}