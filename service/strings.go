package service

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
