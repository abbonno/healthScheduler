package health_scheduler

import (
	"errors"
	"fmt"
	"time"
)

type Fecha struct {
	Dia  int
	Mes  time.Month
	Anio int
}

const anioOrigen = 2025

func NewFecha(dia int, mes time.Month, anio int) (Fecha, error) {
	if anio < anioOrigen {
		return Fecha{}, errors.New("año inválido (por debajo de %d)" + fmt.Sprint(anioOrigen))
	}
	if mes < time.January || mes > time.December {
		return Fecha{}, errors.New("el mes debe estar entre enero y diciembre (de 1 a 12)")
	}
	diasEnMes := DiasDelMes(mes, anio)
	if dia < 1 || dia > diasEnMes {
		return Fecha{}, fmt.Errorf("el día debe estar entre 1 y %d para el mes %s", diasEnMes, mes)
	}
	return Fecha{Dia: dia, Mes: mes, Anio: anio}, nil
}

func EsBisiesto(anio int) bool {
	return anio%4 == 0 && (anio%100 != 0 || anio%400 == 0)
}

func DiasDelMes(mes time.Month, anio int) int {
	switch mes {
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if EsBisiesto(anio) {
			return 29
		}
		return 28
	default:
		return 31
	}
}

func SumarMes(f Fecha) Fecha {
	f.Mes++
	if f.Mes > time.December {
		f.Mes = time.January
		f.Anio++
	}
	return f
}

func SumarDía(f Fecha) Fecha {
	f.Dia++
	if f.Dia > DiasDelMes(f.Mes, f.Anio) {
		f.Dia = 1
		f = SumarMes(f)
	}
	return f
}

func (f Fecha) String() string {
	return fmt.Sprintf("%02d-%02d-%04d", f.Dia, f.Mes, f.Anio)
}

func FechaActual() (Fecha, error) {
	return NewFecha(time.Now().Day(), time.Now().Month(), time.Now().Year())
}
