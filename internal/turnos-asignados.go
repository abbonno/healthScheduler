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

func ObtenerAreaEmpleado(turnosAsignados TurnosAsignados, empleado Empleado, fecha time.Weekday, turno TipoTurno) string {
	for _, t := range turnosAsignados[&empleado] {
		if t.DiaSemana == fecha && t.Nombre == turno {
			return t.Area
		}
	}
	return ""
}
