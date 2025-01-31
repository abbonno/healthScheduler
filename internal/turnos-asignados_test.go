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
			fechaNormalizada := NormalizarFecha(turno.Fecha)
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
	fechaInicio := NormalizarFecha(time.Now())
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

func TestMostrarTurnosEnfermero(t *testing.T) {
	enfermero1 := Empleado("Enfermero1")
	enfermero2 := Empleado("Enfermero2")

	fechaEjemplo := NormalizarFecha(time.Date(2025, time.January, 13, 0, 0, 0, 0, time.UTC))
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
