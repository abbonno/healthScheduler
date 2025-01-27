package health_scheduler

import (
	"strconv"
	"testing"
	"time"
)

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
	ocupacion := make(map[time.Time]map[TipoTurno]map[string]bool)
	for _, turnos := range turnosAsignados {
		for _, turno := range turnos {
			fechaNormalizada := turno.Fecha
			if ocupacion[fechaNormalizada] == nil {
				ocupacion[fechaNormalizada] = make(map[TipoTurno]map[string]bool)
			}
			if ocupacion[fechaNormalizada][turno.Nombre] == nil {
				ocupacion[fechaNormalizada][turno.Nombre] = make(map[string]bool)
			}
			ocupacion[fechaNormalizada][turno.Nombre][turno.Area] = true
		}
	}

	fechaInicio := time.Now()
	for i := 0; i < 365; i++ {
		fecha := fechaInicio.AddDate(0, 0, i)
		for _, tipoTurno := range tiposTurno {
			for _, area := range areas {
				if ocupacion[fecha][tipoTurno][area] == false {
					t.Fatalf("Área %s no ocupada el %s en el turno %s", area, fecha.Format("02-01-2006"), tipoTurno)
				}
			}
		}
	}

}
