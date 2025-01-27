package health_scheduler

import (
	"strconv"
	"testing"
	"time"
)

type Ocupacion map[time.Time]map[TipoTurno]map[string]int

func NewOcupacion(plan TurnosAsignados) Ocupacion {
	ocupacion := make(map[time.Time]map[TipoTurno]map[string]int)
	for _, turnos := range plan {
		for _, turno := range turnos {
			fechaNormalizada := turno.Fecha
			if ocupacion[fechaNormalizada] == nil {
				ocupacion[fechaNormalizada] = make(map[TipoTurno]map[string]int)
			}
			if ocupacion[fechaNormalizada][turno.Nombre] == nil {
				ocupacion[fechaNormalizada][turno.Nombre] = make(map[string]int)
			}
			ocupacion[fechaNormalizada][turno.Nombre][turno.Area]++
		}
	}
	return ocupacion
}

func TestAreaOcupadaDoblemente(t *testing.T) {
	totalEnfermeros := 15
	enfermeros := make([]Empleado, totalEnfermeros)
	for i := 1; i <= totalEnfermeros; i++ {
		e := Empleado("Enfermero" + strconv.Itoa(i))
		enfermeros[i-1] = e
	}
	plan := GenerarPlanAnual(enfermeros)

	listaOcupacion := NewOcupacion(plan)
	for _, turnos := range plan {
		for _, turno := range turnos {
			if listaOcupacion[turno.Fecha][turno.Nombre][turno.Area] > 1 {
				t.Fatalf("Área duplicada: el área %s ya está ocupada el día %v en el turno %v", turno.Area, turno.Fecha, turno.Nombre)
			}
		}
	}
}

func TestAreaSinOcupar(t *testing.T) {
	totalEnfermeros := 15
	enfermeros := make([]Empleado, totalEnfermeros)
	for i := 1; i <= totalEnfermeros; i++ {
		e := Empleado("Enfermero" + strconv.Itoa(i))
		enfermeros[i-1] = e
	}
	tiposTurno := []TipoTurno{Mañana, Tarde, Noche}
	areas := make([]string, totalAreas)
	for i := 1; i <= totalAreas; i++ {
		areas[i-1] = "Area" + strconv.Itoa(i)
	}
	turnosAsignados := GenerarPlanAnual(enfermeros)

	listaOcupacion := NewOcupacion(turnosAsignados)
	fechaInicio := time.Now()
	for i := 0; i < 365; i++ {
		fecha := fechaInicio.AddDate(0, 0, i)
		for _, tipoTurno := range tiposTurno {
			for _, area := range areas {
				if listaOcupacion[fecha][tipoTurno][area] == 0 {
					t.Fatalf("Área %s no ocupada el %s en el turno %s", area, fecha.Format("02-01-2006"), tipoTurno)
				}
			}
		}
	}
}
