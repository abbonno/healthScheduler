package health_scheduler

type Asignacion struct {
	ID         string
	Empleado   Empleado
	Area       string
	Turno      Turno
	EsFlexible bool
}

type CronogramaDiario struct {
	Fecha        string
	Asignaciones []Asignacion
}

func NuevaAsignacion(id string, empleado Empleado, area string, turno Turno, esFlexible bool) *Asignacion {
	return &Asignacion{
		ID:         id,
		Empleado:   empleado,
		Area:       area,
		Turno:      turno,
		EsFlexible: esFlexible,
	}
}
