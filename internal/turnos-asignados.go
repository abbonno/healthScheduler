package health_scheduler

import (
	"errors"
	"fmt"
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
const diasAnio = 365

func GenerarPlanAnual(enfermeros []Empleado) (TurnosAsignados, error) {
	totalEmpleados := len(enfermeros)
	if totalEmpleados < totalAreas {
		return nil, errors.New("Hacen falta, al menos, 10 enfermeros, introducidos: " + fmt.Sprint(totalEmpleados))
	}

	turnosAsignados := make(map[*Empleado][]Turno)
	areas := []string{"Área1", "Área2", "Área3", "Área4", "Área5", "Área6", "Área7", "Área8", "Área9", "Área10"}
	tiposTurno := []TipoTurno{Mañana, Tarde, Noche}
	indiceEnfermero := 0
	fechaActual := time.Now()

	for i := 0; i < diasAnio; i++ {
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
