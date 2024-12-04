package health_scheduler

import (
	"fmt"
	"time"
)

type TipoTurno string

const (
	Mañana TipoTurno = "Mañana"
	Tarde  TipoTurno = "Tarde"
	Noche  TipoTurno = "Noche"
)

var horarios = map[TipoTurno][2]string{
	Mañana: {"08:00", "15:00"},
	Tarde:  {"15:00", "22:00"},
	Noche:  {"22:00", "08:00"},
}

type Turno struct {
	Empleado    string
	NombreTurno TipoTurno
	Area        string
	Fecha       string
	Inicio      string
	Fin         string
}

var RegistroTurnos []Turno // Lista de turnos asignados

func NuevoTurno(empleado string, nombre TipoTurno, area, fecha string) *Turno {

	inicio, fin := horarios[nombre][0], horarios[nombre][1]
	fechaFin := fecha

	if nombre == Noche {
		fechaFin = incrementarDia(fecha) // Día siguiente para turno nocturno
	}

	return &Turno{
		Empleado:    empleado,
		NombreTurno: nombre,
		Area:        area,
		Fecha:       fecha,
		Inicio:      fmt.Sprintf("%s %s", fecha, inicio),
		Fin:         fmt.Sprintf("%s %s", fechaFin, fin),
	}
}

func incrementarDia(fecha string) string {

	const layout = "2006-01-02"
	parsedFecha, err := time.Parse(layout, fecha)
	if err != nil {
		return fecha
	}
	return parsedFecha.AddDate(0, 0, 1).Format(layout)
}
