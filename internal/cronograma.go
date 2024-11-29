package health_scheduler

type RegistroTurnos struct {
	Turnos []Turno // Lista de turnos asignados
}

func NuevoRegistroTurno() *RegistroTurnos {
	return &RegistroTurnos{
		Turnos: []Turno{},
	}
}
