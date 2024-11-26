package health_scheduler

type Empleado struct {
	Nombre         string
	TipoContrato   Contrato
	Turno          Turno
	Area           string
	Disponibilidad []string
}

func NuevoEmpleado(nombre string, tipoContrato Contrato, turno Turno, area string, disponibilidad []string) *Empleado {
	return &Empleado{
		Nombre:         nombre,
		TipoContrato:   tipoContrato,
		Turno:          turno,
		Area:           area,
		Disponibilidad: disponibilidad,
	}
}
