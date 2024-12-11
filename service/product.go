package service

import (
	"fmt"
	"perfume/dao"
	"strings"
)

// Añadir producto
func (s *Service) AddProduct(product dao.Product) {
	s.products = append(s.products, product)
}

// Imprimir el total de productos
func (s Service) ShowTotalProducts() {
	// Contar el total de variantes
	total := 0
	// Contar
	for _, product := range s.products {
		total += len(product.Variants)
	}
	// Imprimir
	fmt.Printf("Total de productos: %d\n", len(s.products))
	fmt.Printf("Total de variantes: %d\n", total)
}

func (s Service) GetProductVariantType(quantity string) (string, dao.ProductVariantType) {
	// Remplazar caracteres
	quantity = strings.ReplaceAll(s.RemoveSpaces(quantity), "|", "")
	// Verificar si hay tamaño
	if len(quantity) > 0 {
		// Verificar si es unidad
		if strings.ToLower(quantity)[len(quantity)-1] == 'u' {
			return quantity[:len(quantity)-1], dao.ProductVariantUnitType
		}
		// Verificar si el tamaño es mayor a dos
		if len(quantity) > 2 {
			runes := []rune(strings.ToLower(quantity))
			// Verificar si es mililitro
			if string(runes[len(runes)-2:]) == "ml" {
				return quantity[:len(quantity)-2], dao.ProductVariantMlType
			}
		}
	}
	// Return
	return "1", dao.ProductVariantUnitType
}

func (s *Service) ValidateProducts() {
	// Iterar productos
	for index, product := range s.products {
		// Verificar si tiene nombre
		if len(product.Name) == 0 {
			fmt.Printf("El producto %d no tiene nombre\n", index)
		}
		// Verificar que tenga marca
		if product.Brand == "" {
			fmt.Printf("El producto %s no tiene marca\n", product.Name)
		}
		// Verificar que tenga categoria
		if product.Category == "" {
			fmt.Printf("El producto %s no tiene categoria\n", product.Name)
		}
		// Verificar que el producto tenga variantes
		if len(product.Variants) > 0 {
			// Iterar variantes
			for vIndex, variant := range product.Variants {
				// Verificar que tenga precio
				if variant.DiscountPrice == "" {
					fmt.Printf("El producto %s no tiene precio\n", product.Name)
				} else {
					// Verificar si existe un precio descuento
					if variant.Price == "" {
						s.products[index].Variants[vIndex].Price = variant.DiscountPrice
					}
				}
				// Verificar que tenga cantidad
				if variant.Quantity == "" {
					fmt.Printf("El producto %s no tiene cantidad\n", product.Name)
				}
				// Verificar si tiene tipo
				if variant.Type == "" {
					fmt.Printf("El producto %s no tiene tipo\n", product.Name)
				}
				// Verificar si tiene fotos
				if len(variant.Photos) == 0 {
					fmt.Printf("El producto %s no tiene fotos\n", product.Name)
				}
			}
		} else {
			fmt.Printf("El producto %s no tiene variantes\n", product.Name)
		}
	}
}
