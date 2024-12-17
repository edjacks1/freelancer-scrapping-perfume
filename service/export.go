package service

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func (s Service) ExportProductsToCsv() error {
	// Crear archivo CSV
	file, err := os.Create("products.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir encabezados
	if err := writer.Write([]string{"Nombre", "Marca", "Categoría", "ID Variación", "Precio", "Descuento", "Fotos", "Activo", "Atributos", "Descripción"}); err != nil {
		return err
	}

	// Escribir datos de productos
	for _, product := range s.products {
		for _, variant := range product.Variants {
			// Determinar estado activo
			isActive := "No"
			if variant.IsActive {
				isActive = "Sí"
			}
			// Atributos dinámicos
			var attributes []string
			if slices.Contains([]string{"ml", "u", "g"}, variant.Type) {
				attributes = append(attributes, "Tamaño:"+variant.Quantity+variant.Type)
			} else {
				fmt.Println(variant.Type)
			}
			// Verificar si es color
			if variant.Color != nil {
				attributes = append(attributes, "Tono:"+variant.Color.Name)
			}
			// Escribir registro
			if err := writer.Write([]string{
				product.Name,
				product.Brand,
				product.Category,
				variant.ID,
				variant.Price,
				variant.DiscountPrice,
				strings.Join(variant.Photos, ";"),
				isActive,
				strings.Join(attributes, ", "),
				variant.Description,
			}); err != nil {
				return err
			}
		}
	}

	log.Println("Archivo CSV creado exitosamente en")
	return nil
}
