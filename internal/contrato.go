package health_scheduler

type Contrato struct {
	Nombre        string
	HorasAnuales  int
	IncluyeNoches bool
}

func NuevoContrato(nombre string, horasAnuales int, incluyeNoches bool) *Contrato {
	return &Contrato{
		Nombre:        nombre,
		HorasAnuales:  horasAnuales,
		IncluyeNoches: incluyeNoches,
	}
}
