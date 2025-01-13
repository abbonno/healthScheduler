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

func NuevoTurno(nombre TipoTurno, area string, fecha time.Time) *Turno {
	return &Turno{
		Nombre: nombre,
		Area:   area,
		Fecha:  fecha,
	}
}

func GenerarPlanSemanal(enfermeros []Empleado) (TurnosAsignados, error) {
	totalEmpleados := len(enfermeros)
	if totalEmpleados < 10 {
		return nil, errors.New("Hacen falta, al menos, 10 enfermeros, introducidos: " + fmt.Sprint(totalEmpleados))
	}

	turnosAsignados := make(map[*Empleado][]Turno)
	areas := []string{"Área1", "Área2", "Área3", "Área4", "Área5", "Área6", "Área7", "Área8", "Área9", "Área10"}
	tiposTurno := []TipoTurno{Mañana, Tarde, Noche}
	indiceEnfermero := 0
	fechaActual := time.Now()

	// Generar turnos para los próximos 365 días
	for i := 0; i < 365; i++ {
		fecha := fechaActual.AddDate(0, 0, i) // Incrementar la fecha día a día

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

func GetTurnosEnfermero(enfermero *Empleado, turnosAsignados TurnosAsignados) []Turno {
	return turnosAsignados[enfermero]
}

func TurnosToString(turnos []Turno) string {
	var result string
	for _, turno := range turnos {
		result += "Fecha: " + turno.Fecha.Format("02-01-2006") + ", Turno: " + string(turno.Nombre) + ", Área: " + turno.Area + "\n"
	}
	return result
}
