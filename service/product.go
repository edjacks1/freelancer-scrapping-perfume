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
