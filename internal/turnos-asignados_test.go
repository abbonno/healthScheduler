package health_scheduler

import (
	"strconv"
	"testing"
	"time"
)

type Ocupacion map[time.Time]map[TipoTurno]map[string]bool

func NewOcupacion(plan TurnosAsignados) Ocupacion {
	ocupacion := make(map[time.Time]map[TipoTurno]map[string]bool)
	for _, turnos := range plan {
		for _, turno := range turnos {
			if ocupacion[turno.Fecha] == nil {
				ocupacion[turno.Fecha] = make(map[TipoTurno]map[string]bool)
			}
			if ocupacion[turno.Fecha][turno.Nombre] == nil {
				ocupacion[turno.Fecha][turno.Nombre] = make(map[string]bool)
			}
		}
	}
	return ocupacion
}

func TestAreaOcupadaDoblemente(t *testing.T) {
	totalEnfermeros := 15
	enfermeros := func() []Empleado {
		base := "Enfermero"
		result := make([]Empleado, totalEnfermeros)
		for i := 1; i <= totalEnfermeros; i++ {
			e := Empleado(base + strconv.Itoa(i))
			result[i-1] = e
		}
		return result
	}()

	plan, err := GenerarPlanAnual(enfermeros)
	if err != nil {
		t.Fatalf("Error generando el plan anual: %v", err)
	}

	listaOcupacion := NewOcupacion(plan)

	for _, turnos := range plan {
		for _, turno := range turnos {
			if listaOcupacion[turno.Fecha][turno.Nombre][turno.Area] {
				t.Errorf("Área duplicada: el área %s ya está ocupada el día %v en el turno %v", turno.Area, turno.Fecha, turno.Nombre)
			}

			listaOcupacion[turno.Fecha][turno.Nombre][turno.Area] = true
		}
	}
}

func TestAreaSinOcupar(t *testing.T) {
	totalEnfermeros := 15
	enfermeros := func() []Empleado {
		base := "Enfermero"
		result := make([]Empleado, totalEnfermeros)
		for i := 1; i <= totalEnfermeros; i++ {
			e := Empleado(base + strconv.Itoa(i))
			result[i-1] = e
		}
		return result
	}()

	turnosAsignados, err := GenerarPlanAnual(enfermeros)
	if err != nil {
		t.Fatalf("Error al generar el plan anual: %v", err)
	}

	listaOcupacion := NewOcupacion(turnosAsignados)

	for _, turnos := range turnosAsignados {
		for _, turno := range turnos {
			listaOcupacion[turno.Fecha][turno.Nombre][turno.Area] = true
		}
	}

	for dia, turnos := range listaOcupacion {
		for turno, areas := range turnos {
			for area, ocupada := range areas {
				if !ocupada {
					t.Errorf("Área %s no está cubierta en el turno %v del día %v", area, turno, dia)
				}
			}
		}
	}

}

func CrearFecha(año int, mes time.Month, día int) time.Time {
	return time.Date(año, mes, día, 0, 0, 0, 0, time.UTC)
}

func TestMostrarTurnosEnfermero(t *testing.T) {
	enfermero1 := Empleado("Enfermero1")
	enfermero2 := Empleado("Enfermero2")

	fechaEjemplo := CrearFecha(2025, time.January, 13)

	planSemanal := TurnosAsignados{
		&enfermero1: {
			{Fecha: fechaEjemplo, Area: "Área1", Nombre: Mañana},
			{Fecha: fechaEjemplo.AddDate(0, 0, 1), Area: "Área2", Nombre: Tarde},
		},
		&enfermero2: {
			{Fecha: fechaEjemplo, Area: "Área2", Nombre: Noche},
			{Fecha: fechaEjemplo.AddDate(0, 0, 2), Area: "Área3", Nombre: Mañana},
		},
	}

	resultado := MostrarTurnosEnfermero(&enfermero1, planSemanal)

	esperado := "Fecha: 13-01-2025, Turno: Mañana, Área: Área1\nFecha: 14-01-2025, Turno: Tarde, Área: Área2\n"

	if resultado != esperado {
		t.Errorf("Resultado inesperado: got %q \n want %q", resultado, esperado)
	}
}
