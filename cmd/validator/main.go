package main

import (
	"fmt"
	"perfume/service"
)

func main() {
	// Inicializar servicio
	svc := service.InitService()
	// Cargar productos
	if err := svc.LoadFile(); err == nil {
		svc.ValidateProducts()
		//
		// svc.ShowAllAttributes()
		// wcService := woocommerce.InitService()
		// Verificar atributos
		// wcService.VerifyAttributes(svc.GetUniqueAttributes())
		// Exportar
		fmt.Println(svc.ExportProductsToCsv())
	}
}
