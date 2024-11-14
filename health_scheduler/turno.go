package health_scheduler

type Turno struct {
	Nombre TipoTurno
	Inicio string
	Fin    string
}

func NuevoTurno(nombre TipoTurno) *Turno {

	var inicio, fin string

	switch nombre {
	case Mañana:
		inicio = "08:00"
		fin = "15:00"
	case Tarde:
		inicio = "15:00"
		fin = "22:00"
	case Noche:
		inicio = "22:00"
		fin = "08:00"
	}

	return &Turno{
		Nombre: nombre,
		Inicio: inicio,
		Fin:    fin,
	}

}
