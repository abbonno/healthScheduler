package health_scheduler

type Empleado struct {
	ID           string  // Identificador único (puede ser un UUID o similar)
	Nombre       string  // Nombre del enfermero
	TipoContrato string  // Tipo de contrato (puede ser "Tiempo completo", "Medio tiempo", etc.)
	Turnos       []Turno // Lista de turnos asignados al enfermero
}

func NuevoEnfermero(id, nombre, tipoContrato string) *Empleado {
	return &Empleado{
		ID:           id,
		Nombre:       nombre,
		TipoContrato: tipoContrato,
		Turnos:       []Turno{}, // Inicialmente sin turnos asignados
	}
}
