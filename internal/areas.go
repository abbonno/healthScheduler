package health_scheduler

type Area struct {
	Nombre               string
	EnfermerosNecesarios int
}

var (
	Intermedia1 = Area{Nombre: "Intermedia 1", EnfermerosNecesarios: 1}
	Intermedia2 = Area{Nombre: "Intermedia 2", EnfermerosNecesarios: 1}
	Intermedia3 = Area{Nombre: "Intermedia 3", EnfermerosNecesarios: 1}
	Intermedia4 = Area{Nombre: "Intermedia 4", EnfermerosNecesarios: 1}
	Urgencias1  = Area{Nombre: "Urgencias 1", EnfermerosNecesarios: 1}
	Urgencias2  = Area{Nombre: "Urgencias 2", EnfermerosNecesarios: 1}
	Planta1     = Area{Nombre: "Planta 1", EnfermerosNecesarios: 1}
	Planta2     = Area{Nombre: "Planta 2", EnfermerosNecesarios: 1}
	Sustitutos  = Area{Nombre: "Sustitutos", EnfermerosNecesarios: 2}
)
