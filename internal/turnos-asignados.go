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
	Fecha  time.Time
}

type Empleado string

type TurnosAsignados map[*Empleado][]Turno

const diasAnio = 365
