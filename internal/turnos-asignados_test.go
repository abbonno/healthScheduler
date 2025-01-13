package health_scheduler

import (
	"testing"
	"time"
)

func TestGenerarPlanSemanal(t *testing.T) {
	enfermeros := []Empleado{
		"Enfermero1", "Enfermero2", "Enfermero3", "Enfermero4", "Enfermero5",
		"Enfermero6", "Enfermero7", "Enfermero8", "Enfermero9", "Enfermero10",
		"Enfermero11", "Enfermero12", "Enfermero13", "Enfermero14", "Enfermero15",
	}

	plan, err := GenerarPlanSemanal(enfermeros)
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

func TestGetTurnosEnfermero(t *testing.T) {
	enfermero1 := Empleado("Enfermero1")
	enfermero2 := Empleado("Enfermero2")

	planSemanal := TurnosAsignados{
		&enfermero1: {
			{Fecha: time.Now(), Area: "Área1", Nombre: Mañana},
			{Fecha: time.Now().AddDate(0, 0, 1), Area: "Área2", Nombre: Tarde},
		},
		&enfermero2: {
			{Fecha: time.Now(), Area: "Área2", Nombre: Noche},
			{Fecha: time.Now().AddDate(0, 0, 1), Area: "Área3", Nombre: Mañana},
		},
	}

	turnos := GetTurnosEnfermero(&enfermero1, planSemanal)

	if len(turnos) == 0 {
		t.Errorf("El enfermero %s no tiene turnos asignados", enfermero1)
	}

	for _, turno := range turnos {
		if turno.Area == "" || turno.Nombre == "" || turno.Fecha.String() == "" {
			t.Errorf("Turno inválido encontrado: %+v", turno)
		}
	}

	enfermeroNoAsignado := Empleado("EnfermeroSinTurnos")
	turnosSinAsignar := GetTurnosEnfermero(&enfermeroNoAsignado, planSemanal)

	if len(turnosSinAsignar) != 0 {
		t.Errorf("El enfermero %s tiene turnos asignados, pero no debería", enfermeroNoAsignado)
	}
}

func TestTurnosToString(t *testing.T) {
	fecha := time.Date(2025, time.January, 13, 0, 0, 0, 0, time.UTC)
	turnos := []Turno{
		{
			Nombre: Mañana,
			Area:   "A1",
			Fecha:  fecha,
		},
	}

	resultado := TurnosToString(turnos)

	esperado := "Fecha: 13-01-2025, Turno: Mañana, Área: A1\n"

	if resultado != esperado {
		t.Errorf("Resultado inesperado: got %v, want %v", resultado, esperado)
	}
}
