package health_scheduler

import (
	"testing"
	"time"
)

func TestAreaOcupadaDoblemente(t *testing.T) {
	enfermeros := []Empleado{
		"Enfermero1", "Enfermero2", "Enfermero3", "Enfermero4", "Enfermero5",
		"Enfermero6", "Enfermero7", "Enfermero8", "Enfermero9", "Enfermero10",
		"Enfermero11", "Enfermero12", "Enfermero13", "Enfermero14", "Enfermero15",
	}

	plan, err := GenerarPlanAnual(enfermeros)
	if err != nil {
		t.Fatalf("Error generando el plan anual: %v", err)
	}

	turnosPorDia := make(map[time.Time]map[TipoTurno]map[string]bool)

	for _, turnos := range plan {
		for _, turno := range turnos {
			if turnosPorDia[turno.Fecha] == nil {
				turnosPorDia[turno.Fecha] = make(map[TipoTurno]map[string]bool)
			}
			if turnosPorDia[turno.Fecha][turno.Nombre] == nil {
				turnosPorDia[turno.Fecha][turno.Nombre] = make(map[string]bool)
			}
			if turnosPorDia[turno.Fecha][turno.Nombre][turno.Area] {
				t.Errorf("Área duplicada: el área %s ya está ocupada el día %v en el turno %v", turno.Area, turno.Fecha, turno.Nombre)
			}

			turnosPorDia[turno.Fecha][turno.Nombre][turno.Area] = true
		}
	}
}
