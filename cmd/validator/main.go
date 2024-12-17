package main

import "perfume/service"

func main() {
	// Inicializar servicio
	svc := service.InitService()
	// Cargar productos
	if err := svc.LoadFile(); err == nil {
		svc.ValidateProducts()
		// Exportar
		svc.ExportProductsToCsv()
	}
}
