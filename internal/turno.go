package health_scheduler

import (
	"time"
)

type TipoTurno string

const (
	Mañana TipoTurno = "Mañana"
	Tarde  TipoTurno = "Tarde"
	Noche  TipoTurno = "Noche"
)

type Turno struct {
	Nombre TipoTurno
	Area   string
	Fecha  time.Weekday
}

func NuevoTurno(nombre TipoTurno, area string, fecha time.Weekday) *Turno {
	return &Turno{
		Nombre: nombre,
		Area:   area,
		Fecha:  fecha,
	}
}
