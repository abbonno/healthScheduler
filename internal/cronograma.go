package health_scheduler

type DiaSemana string

const (
	Lunes     DiaSemana = "Lunes"
	Martes    DiaSemana = "Martes"
	Miercoles DiaSemana = "Miercoles"
	Jueves    DiaSemana = "Jueves"
	Viernes   DiaSemana = "Viernes"
)

type CronogramaSemanal struct {
	Dia          DiaSemana
	Asignaciones []Asignacion
}

// Método para crear un nuevo cronograma para un día específico
func NuevoCronograma(dia DiaSemana) *CronogramaSemanal {
	return &CronogramaSemanal{
		Dia:          dia,
		Asignaciones: []Asignacion{}, // Inicializa la lista vacía
	}
}

// Método para añadir una asignación al cronograma
func (c *CronogramaSemanal) AgregarAsignacion(asignacion Asignacion) {
	c.Asignaciones = append(c.Asignaciones, asignacion)
}
