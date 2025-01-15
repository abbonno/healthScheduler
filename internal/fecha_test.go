package health_scheduler

import (
	"testing"
	"time"
)

func TestEsBisiesto(t *testing.T) {
	aniosPrueba := []struct {
		anio       int
		esBisiesto bool
	}{
		{2024, true},
		{2025, false},
		{2026, false},
		{2027, false},
		{2028, true},
	}
	for _, anio := range aniosPrueba {
		if resultado := EsBisiesto(anio.anio); resultado != anio.esBisiesto {
			t.Errorf("EsBisiesto(%v) = %v; se esperaba %v", anio.anio, resultado, anio.esBisiesto)
		}
	}
}

func TestDiasDelMes(t *testing.T) {
	fechasPrueba := []struct {
		mes      time.Month
		anio     int
		esperado int
	}{
		{time.January, 2025, 31},
		{time.February, 2025, 28},
		{time.February, 2024, 29},
		{time.March, 2025, 31},
		{time.April, 2025, 30},
		{time.May, 2025, 31},
		{time.June, 2025, 30},
		{time.July, 2025, 31},
		{time.August, 2025, 31},
		{time.September, 2025, 30},
		{time.October, 2025, 31},
		{time.November, 2025, 30},
		{time.December, 2025, 31},
	}
	for _, fecha := range fechasPrueba {
		if resultado := DiasDelMes(fecha.mes, fecha.anio); resultado != fecha.esperado {
			t.Errorf("DiasDelMes(%v, %v) = %v; se esperaba %v", fecha.mes, fecha.anio, resultado, fecha.esperado)
		}
	}
}

func TestSumarMes(t *testing.T) {
	fecha := Fecha{Dia: 1, Mes: time.January, Anio: 2025}
	fecha = SumarMes(fecha)
	if fecha.Mes != time.February {
		t.Errorf("SumarMes() = %v; se esperaba %v", fecha.Mes, time.February)
	}
	fecha.Mes = time.December
	fecha = SumarMes(fecha)
	if fecha.Mes != time.January {
		t.Errorf("SumarMes() = %v; se esperaba %v", fecha.Mes, time.January)
	}
}

func TestSumarDía(t *testing.T) {
	fecha := Fecha{Dia: 1, Mes: time.January, Anio: 2025}
	fecha = SumarDía(fecha)
	if fecha.Dia != 2 {
		t.Errorf("SumarDía() = %v; se esperaba %v", fecha.Dia, 2)
	}
	fecha.Dia = 31
	fecha = SumarDía(fecha)
	if fecha.Dia != 1 {
		t.Errorf("SumarDía() = %v; se esperaba %v", fecha.Dia, 1)
	}
	if fecha.Mes != time.February {
		t.Errorf("SumarDía() = %v; se esperaba %v", fecha.Mes, time.February)
	}
}

func TestFechaActual(t *testing.T) {
	fecha, _ := FechaActual()
	ahora := time.Now()
	if fecha.Dia != ahora.Day() || fecha.Mes != ahora.Month() || fecha.Anio != ahora.Year() {
		t.Errorf("FechaActual() = %v; se esperaba %v", fecha, ahora)
	}
}
