package health_scheduler

type RegistroTurnos []Turno // Lista de turnos asignados

func NuevoRegistroTurno() *RegistroTurnos {
	rt := RegistroTurnos{}
	return &rt
}
