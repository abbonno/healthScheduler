package health_scheduler

import (
	"testing"
	"time"
)

func TestNuevoTurno(t *testing.T) {
	// Datos de entrada
	nombre := Mañana
	area := "Urgencias"
	fecha := time.Monday

	// Llamada al constructor
	turno := NuevoTurno(nombre, area, fecha)

	// Aserciones
	if turno.Nombre != nombre {
		t.Errorf("Error: Nombre del turno = %s; esperado = %s", turno.Nombre, nombre)
	}

	if turno.Area != area {
		t.Errorf("Error: Área del turno = %s; esperado = %s", turno.Area, area)
	}

	if turno.Fecha != fecha {
		t.Errorf("Error: Fecha del turno = %v; esperado = %v", turno.Fecha, fecha)
	}
}
