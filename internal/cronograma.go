package health_scheduler

import (
	"fmt"
)

type RegistroTurnos struct {
	Turnos []Turno // Lista de turnos asignados
}

func NuevoRegistroTurno() *RegistroTurnos {
	return &RegistroTurnos{
		Turnos: []Turno{},
	}
}

func (e *RegistroTurnos) AgregarTurno(turno Turno) {
	e.Turnos = append(e.Turnos, turno)
}

func (e *RegistroTurnos) ListarTurnos() {
	for _, turno := range e.Turnos {
		fmt.Printf("Empleado: %s, Turno: %s, Fecha: %s, Inicio: %s, Fin: %s, Área: %s\n",
			turno.Empleado.Nombre,
			turno.Nombre,
			turno.Fecha,
			turno.Inicio,
			turno.Fin,
			turno.Area)
	}
}

func (e *RegistroTurnos) ObtenerTurnoPorEmpleado(nombreEmpleado string) *Turno {
	for _, turno := range e.Turnos {
		if turno.Empleado.Nombre == nombreEmpleado {
			return &turno
		}
	}
	return nil // Retorna nil si no se encuentra el turno
}
