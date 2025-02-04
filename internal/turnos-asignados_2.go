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
	Nombre    TipoTurno
	Area      string
	DiaSemana time.Weekday
}

type Empleado string

type TurnosAsignados map[*Empleado][]Turno
