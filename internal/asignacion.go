package health_scheduler

type Asignacion struct {
	Empleado   Empleado
	EsFlexible bool
}

func NuevaAsignacion(empleado Empleado, esFlexible bool) *Asignacion {
	return &Asignacion{
		Empleado:   empleado,
		EsFlexible: esFlexible,
	}
}
