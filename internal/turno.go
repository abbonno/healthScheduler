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

type HorarioMañana string

const (
	InicioMañana HorarioMañana = "08:00"
	FinMañana    HorarioMañana = "15:00"
)

type HorarioTarde string

const (
	InicioTarde HorarioTarde = "15:00"
	FinTarde    HorarioTarde = "22:00"
)

type HorarioNoche string

const (
	InicioNoche HorarioNoche = "22:00"
	FinNoche    HorarioNoche = "08:00"
)

type Turno struct {
	Empleado Empleado
	Nombre   TipoTurno
	Area     string
	Fecha    string
	Inicio   string
	Fin      string
}

func NuevoTurno(empleado Empleado, nombre TipoTurno, area, fecha string) *Turno {
	var inicio, fin, fechaFin string

	switch nombre {
	case Mañana:
		inicio = string(InicioMañana)
		fin = string(FinMañana)
		fechaFin = fecha // Mismo día
	case Tarde:
		inicio = string(InicioTarde)
		fin = string(FinTarde)
		fechaFin = fecha // Mismo día
	case Noche:
		inicio = string(InicioNoche)
		fin = string(FinNoche)
		fechaFin = incrementarDia(fecha) // Día siguiente
	default:
		inicio, fin, fechaFin = "", "", fecha
	}

	return &Turno{
		Empleado: empleado,
		Nombre:   nombre,
		Area:     area,
		Fecha:    fecha,
		Inicio:   fmt.Sprintf("%s %s", fecha, inicio),
		Fin:      fmt.Sprintf("%s %s", fechaFin, fin),
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
