package service

import (
	"fmt"
	"perfume/dao"
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
