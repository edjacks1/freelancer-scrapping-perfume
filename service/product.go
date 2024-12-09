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
	fmt.Printf("Total de productos: %d\n", len(s.products))
}
