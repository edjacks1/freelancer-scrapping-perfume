package service

import (
	"encoding/csv"
	"log"
	"os"
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
	if err := writer.Write([]string{
		"sku",
		"Nombre",
		"Descripción",
		"Precio rebajado",
		"Precio normal",
		"Categorías",
		"Nombre del atributo 1",
		"Valor(es) del atributo 1",
		"Atributo visible 1",
		"Atributo global 1",
		"Imágenes", // Nueva columna para imágenes
		"Marca",
	}); err != nil {
		return err
	}

	// Escribir datos de productos
	for _, product := range s.products {
		// Escribir una fila para cada variante del producto
		for _, variant := range product.Variants {
			// Inicializar las listas de atributos y valores
			var attributeNames string
			var attributeValues string
			var attributeVisible string
			var attributeGlobal string
			// Recopilar el SKU de la variante (ID del proveedor)
			variantSKU := variant.ID

			// Recopilar precios
			regularPrice := variant.Price
			discountPrice := variant.DiscountPrice

			// Recopilar categorías (puede ser un solo valor o varias, dependiendo de cómo se maneje el producto)
			category := strings.ReplaceAll(product.Category, "/", ">")

			// Atributos del producto (por ejemplo, "Tamaño" o "Color")
			if variant.Color != nil {
				attributeNames = "Tono"
				attributeValues = variant.Color.Name
				attributeVisible = "1" // Este atributo será visible
				attributeGlobal = "1"  // Esto depende de si el atributo es global
			} else {
				attributeNames = "Tamaño"
				attributeValues = variant.Quantity + "" + variant.Type
				attributeVisible = "1" // Este atributo será visible
				attributeGlobal = "1"  // Esto depende de si el atributo es global
			}
			// Descripción de la variante
			description := variant.Description

			// Escribir una fila para la variante como producto independiente
			if err := writer.Write([]string{
				variantSKU,
				product.Name,                      // Nombre del producto
				description,                       // Descripción de la variante
				discountPrice,                     // Precio rebajado de la variante
				regularPrice,                      // Precio normal de la variante
				category,                          // Categorías del producto
				attributeNames,                    // Nombre del atributo 1
				attributeValues,                   // Valor(es) del atributo 1
				attributeVisible,                  // Atributo visible 1
				attributeGlobal,                   // Atributo global 1
				strings.Join(variant.Photos, ","), // Imágenes (nueva columna)
				product.Brand,
			}); err != nil {
				return err
			}
		}
	}

	log.Println("Archivo CSV creado exitosamente")
	return nil
}
