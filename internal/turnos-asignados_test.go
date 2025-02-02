package health_scheduler

import (
	"testing"
	"time"
)

func TestObtenerAreaEmpleado(t *testing.T) {
	turnosAsignados := make(TurnosAsignados)
	e1 := Empleado("E1")
	e2 := Empleado("E2")
	e3 := Empleado("E3")
	turnosAsignados[&e1] = []Turno{{Nombre: "Mañana", Area: "A1", Fecha: time.Monday}}
	turnosAsignados[&e2] = []Turno{{Nombre: "Mañana", Area: "A2", Fecha: time.Monday}}
	turnosAsignados[&e3] = []Turno{{Nombre: "Mañana", Area: "A3", Fecha: time.Monday}}

	empleado := "E2"
	fecha := time.Monday
	turno := "Mañana"
	areaEsperada := "A2"

	obtenido := ObtenerAreaEmpleado(turnosAsignados, empleado, fecha, turno)
	if obtenido != areaEsperada {
		t.Errorf("obtenido: %s, esperado: %s", obtenido, areaEsperada)
	}
}
