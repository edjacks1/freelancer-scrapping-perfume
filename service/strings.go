package service

import (
	"regexp"
	"strings"
)

func (s Service) RemoveDuplicates(strings []string) []string {
	// Usar un mapa para almacenar los elementos únicos
	unique := make(map[string]struct{})

	// Iterar sobre el slice y agregar los elementos al mapa
	for _, str := range strings {
		unique[str] = struct{}{}
	}

	// Crear un slice para los elementos únicos
	result := make([]string, 0, len(unique))
	for str := range unique {
		result = append(result, str)
	}

	return result
}

func (s Service) RemoveSpaces(value string) string {
	return strings.ReplaceAll(value, " ", "")
}

func (s Service) GetDigitsFromString(value string) string {
	return regexp.MustCompile(`[^0-9.-]+`).ReplaceAllString(strings.ReplaceAll(value, ",", "."), "")
}
