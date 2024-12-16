package service

import (
	"fmt"
	"time"
)

func (s Service) CalculateDuration(start, end time.Time) string {
	// Calcula la diferencia entre las dos fechas
	duration := end.Sub(start)

	// Si el tiempo es negativo, significa que la fecha final es anterior a la inicial
	if duration < 0 {
		return "El tiempo final es anterior al tiempo inicial."
	}

	// Obtiene las horas, minutos y segundos
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	// Formatea la salida dependiendo de la duración
	switch {
	case hours == 0 && minutes == 0:
		// Menor a un minuto
		return fmt.Sprintf("Total: %d segundos", seconds)
	case hours == 0:
		// Menor a una hora
		return fmt.Sprintf("Total: %d minutos y %d segundos", minutes, seconds)
	default:
		// Una hora o más
		return fmt.Sprintf("Total: %d horas, %d minutos y %d segundos", hours, minutes, seconds)
	}
}
