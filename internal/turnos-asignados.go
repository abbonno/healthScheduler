package health_scheduler

import (
	"strconv"
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

const totalAreas = 10
const diasPlan = 365

func NormalizarFecha(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func GenerarPlanAnual(enfermeros []Empleado) TurnosAsignados {
	turnosAsignados := make(map[*Empleado][]Turno)
	fechaActual := NormalizarFecha(time.Now())
	tiposTurno := []TipoTurno{Mañana, Tarde, Noche}
	areas := make([]string, totalAreas)
	for i := 1; i <= totalAreas; i++ {
		areas[i-1] = "Area" + strconv.Itoa(i)
	}
	indiceEnfermero := 0

	for i := 0; i < diasPlan; i++ {
		fecha := fechaActual.AddDate(0, 0, i)
		for _, tipoTurno := range tiposTurno {
			for _, area := range areas {
				turno := Turno{Fecha: fecha, Area: area, Nombre: tipoTurno}
				turnosAsignados[&enfermeros[indiceEnfermero]] = append(turnosAsignados[&enfermeros[indiceEnfermero]], turno)
				indiceEnfermero = (indiceEnfermero + 1) % len(enfermeros)
			}
		}
	}
	return turnosAsignados
}

func MostrarTurnosEnfermero(enfermero *Empleado, plan TurnosAsignados) string {
	turnos := plan[enfermero]
	var resultado string
	for _, turno := range turnos {
		resultado += "Fecha: " + turno.Fecha.Format("02-01-2006") + ", Turno: " + string(turno.Nombre) + ", Área: " + turno.Area + "\n"
	}
	return resultado
}
