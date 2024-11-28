package health_scheduler

type Empleado struct {
	Nombre       string
	TipoContrato Contrato
}

func NuevoEmpleado(nombre string, tipoContrato Contrato) *Empleado {
	return &Empleado{
		Nombre:       nombre,
		TipoContrato: tipoContrato,
	}
}
