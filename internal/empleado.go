package health_scheduler

type Empleado struct {
	Nombre         string
	TipoContrato   Contrato
	Turno          Turno
	Area           Area
	Disponibilidad []string
}

func NuevoEmpleado(nombre string, tipoContrato Contrato, turno Turno, area Area, disponibilidad []string) *Empleado {
	return &Empleado{
		Nombre:         nombre,
		TipoContrato:   tipoContrato,
		Turno:          turno,
		Area:           area,
		Disponibilidad: disponibilidad,
	}
}
