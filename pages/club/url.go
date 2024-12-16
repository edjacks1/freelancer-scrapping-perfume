package pages_club

import (
	"fmt"
	"strings"
)

// Obtener categoria de la url
func (p Page) GetCategoryFromUrl(url string) (string, error) {
	// Divide la URL en segmentos usando "/" como separador
	parts := strings.Split(strings.TrimSuffix(url, "/"), "/")
	// Comprueba que hay al menos dos segmentos
	if len(parts) < 2 {
		return "", fmt.Errorf("No se puede extraer el penúltimo valor %s", url)
	}
	// Retorna el penúltimo segmento
	return parts[len(parts)-2], nil
}
