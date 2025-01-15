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

func TestAreaSinOcupar(t *testing.T) {
	enfermeros := []Empleado{
		"Enfermero1", "Enfermero2", "Enfermero3", "Enfermero4", "Enfermero5",
		"Enfermero6", "Enfermero7", "Enfermero8", "Enfermero9", "Enfermero10",
		"Enfermero11", "Enfermero12", "Enfermero13", "Enfermero14", "Enfermero15",
	}

	turnosAsignados, err := GenerarPlanAnual(enfermeros)
	if err != nil {
		t.Fatalf("Error al generar el plan anual: %v", err)
	}

	turnosPorDia := make(map[time.Time]map[TipoTurno]map[string]bool)

	for _, turnos := range turnosAsignados {
		for _, turno := range turnos {
			if turnosPorDia[turno.Fecha] == nil {
				turnosPorDia[turno.Fecha] = make(map[TipoTurno]map[string]bool)
			}
			if turnosPorDia[turno.Fecha][turno.Nombre] == nil {
				turnosPorDia[turno.Fecha][turno.Nombre] = make(map[string]bool)
			}

			turnosPorDia[turno.Fecha][turno.Nombre][turno.Area] = true
		}
	}

	for dia, turnos := range turnosPorDia {
		for turno, areas := range turnos {
			if len(areas) < totalAreas {
				t.Errorf("Faltan áreas en el turno %v del día %v: solo hay %d áreas cubiertas", turno, dia, len(areas))
			}
			for area, ocupada := range areas {
				if !ocupada {
					t.Errorf("Área %s no está cubierta en el turno %v del día %v", area, turno, dia)
				}
			}
		}
	}

}

func TestMostrarTurnosEnfermero(t *testing.T) {
	enfermero1 := Empleado("Enfermero1")
	enfermero2 := Empleado("Enfermero2")

	fecha := time.Date(2025, time.January, 13, 0, 0, 0, 0, time.UTC)

	planSemanal := TurnosAsignados{
		&enfermero1: {
			{Fecha: fecha, Area: "Área1", Nombre: Mañana},
			{Fecha: fecha.AddDate(0, 0, 1), Area: "Área2", Nombre: Tarde},
		},
		&enfermero2: {
			{Fecha: fecha, Area: "Área2", Nombre: Noche},
			{Fecha: fecha.AddDate(0, 0, 1), Area: "Área3", Nombre: Mañana},
		},
	}

	resultado := MostrarTurnosEnfermero(&enfermero1, planSemanal)

	esperado := "Fecha: 13-01-2025, Turno: Mañana, Área: Área1\nFecha: 14-01-2025, Turno: Tarde, Área: Área2\n"

	if resultado != esperado {
		t.Errorf("Resultado inesperado: got %q \n want %q", resultado, esperado)
	}
}
