package domain

import (
	"strings";
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

type Registro struct {
	ID     int       `json:"id"`
	Data   Date      `json:"data"`
	TipoID   int       `json:"tipo_id"`
}