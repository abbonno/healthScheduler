package health_scheduler

import (
	"errors"
	"fmt"
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

func GenerarPlanAnual(enfermeros []Empleado) (TurnosAsignados, error) {
	totalEmpleados := len(enfermeros)
	if totalEmpleados < totalAreas {
		return nil, errors.New("Hacen falta, al menos, 10 enfermeros, introducidos: " + fmt.Sprint(totalEmpleados))
	}

	turnosAsignados := make(map[*Empleado][]Turno)

	fechaActual := time.Now()
	tiposTurno := []TipoTurno{Mañana, Tarde, Noche}
	areas := func() []string {
		base := "Área"
		result := make([]string, totalAreas)
		for i := 1; i <= totalAreas; i++ {
			result[i-1] = base + strconv.Itoa(i)
		}
		return result
	}()
	indiceEnfermero := 0

	for i := 0; i < diasPlan; i++ {
		fecha := fechaActual.AddDate(0, 0, i)
		for _, tipoTurno := range tiposTurno {
			for _, area := range areas {
				empleado := &enfermeros[indiceEnfermero]

				turno := Turno{
					Fecha:  fecha,
					Area:   area,
					Nombre: tipoTurno,
				}

				if _, ok := turnosAsignados[empleado]; !ok {
					turnosAsignados[empleado] = []Turno{}
				}

				turnosAsignados[empleado] = append(turnosAsignados[empleado], turno)
				indiceEnfermero = (indiceEnfermero + 1) % totalEmpleados
			}
		}
	}

	return turnosAsignados, nil
}

func MostrarTurnosEnfermero(enfermero *Empleado, plan TurnosAsignados) string {
	turnos := plan[enfermero]
	var resultado string
	for _, turno := range turnos {
		resultado += "Fecha: " + turno.Fecha.Format("02-01-2006") + ", Turno: " + string(turno.Nombre) + ", Área: " + turno.Area + "\n"
	}
	return resultado
}
