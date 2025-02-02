package health_scheduler

import (
	"testing"
	"time"
)

func TestObtenerAreaEmpleado(t *testing.T) {
	turnosAsignados := make(TurnosAsignados)
	fecha := time.Monday
	turno := Mañana
	e1 := Empleado("E1")
	e2 := Empleado("E2")
	e3 := Empleado("E3")
	turnosAsignados[&e1] = []Turno{{Nombre: turno, Area: "A1", Fecha: fecha}}
	turnosAsignados[&e2] = []Turno{{Nombre: turno, Area: "A2", Fecha: fecha}}
	turnosAsignados[&e3] = []Turno{{Nombre: turno, Area: "A3", Fecha: fecha}}

	areaEsperada := "A2"

	obtenido := ObtenerAreaEmpleado(turnosAsignados, e2, fecha, turno)
	if obtenido != areaEsperada {
		t.Errorf("obtenido: %s, esperado: %s", obtenido, areaEsperada)
	}
}
