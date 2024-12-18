package service

import (
	"fmt"
	"html"
	"perfume/constants"
	"perfume/domain/dao"
	"regexp"
	"slices"
	"strings"
)

// Añadir producto
func (s *Service) AddWrongUrls(url string) {
	s.wrongUrls = append(s.wrongUrls, url)
}

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
	fmt.Printf("Total de variantes: %d\n\n", total)
	// Verificar si existen urls malas
	if len(s.wrongUrls) > 0 {
		fmt.Printf("Urls que tuvieron problemas: %d\n\n", total)
		// Iterar
		for _, url := range s.wrongUrls {
			fmt.Println(url)
		}
	}
}

func (s Service) GetProductVariantType(quantity string) (string, string) {
	// Remplazar caracteres
	quantity = strings.ReplaceAll(s.RemoveSpaces(quantity), "|", "")
	// Verificar si hay tamaño
	if len(quantity) > 0 {
		// Verificar si es unidad
		if strings.ToLower(quantity)[len(quantity)-1] == 'u' {
			return quantity[:len(quantity)-1], string(dao.ProductVariantUnitType)
		}
		// Verificar si el tamaño es mayor a dos
		if len(quantity) > 2 {
			runes := []rune(strings.ToLower(quantity))
			// Verificar si es mililitro
			if string(runes[len(runes)-2:]) == "ml" {
				return quantity[:len(quantity)-2], string(dao.ProductVariantMlType)
			}
		}
	}
	// Return
	return "1", string(dao.ProductVariantUnitType)
}

func (s *Service) ValidateProducts() {
	// Iterar productos
	for index, product := range s.products {
		// Verificar si tiene nombre
		if len(product.Name) == 0 {
			fmt.Printf("El producto %d no tiene nombre\n", index)
		} else {
			s.products[index].Name = strings.TrimSpace(html.UnescapeString(product.Name))
		}
		// Verificar que tenga marca
		if product.Brand == "" {
			fmt.Printf("El producto %s no tiene marca\n", product.Name)
		} else {
			s.products[index].Brand = strings.TrimSpace(html.UnescapeString(product.Brand))
		}
		// Verificar que tenga categoria
		if product.Category == "" {
			fmt.Printf("El producto %s no tiene categoria\n", product.Name)
		} else {
			s.products[index].Category = strings.TrimSpace(html.UnescapeString(product.Category))
		}
		// Verificar que el producto tenga variantes
		if len(product.Variants) > 0 {
			// Iterar variantes
			for vIndex, variant := range product.Variants {
				// Actualizar descripcion
				s.products[index].Variants[vIndex].Description = strings.TrimSpace(html.UnescapeString(variant.Description))
				// Verificar que tenga precio
				if variant.DiscountPrice == "" || variant.DiscountPrice == "0" {
					fmt.Printf("El producto %s no tiene precio\n", product.Name)
				} else {
					// Verificar si existe un precio descuento
					if variant.Price == "" || variant.Price == "0" {
						s.products[index].Variants[vIndex].Price = variant.DiscountPrice
					}
				}
				// Verificar que tenga cantidad
				if variant.Quantity == "" {
					fmt.Printf("El producto %s no tiene cantidad\n", product.Name)
				} else {
					// Remplazar comas por puntos
					variant.Quantity = strings.ReplaceAll(variant.Quantity, ",", ".")
					s.products[index].Variants[vIndex].Quantity = variant.Quantity
					// Verificar si es un dato valido
					if !regexp.MustCompile(`^\d+(\.\d+)?$`).MatchString(variant.Quantity) {
						fmt.Printf("El producto %s no tiene formato de cantidad incorrecto %s\n", product.Name, variant.Quantity)
					}
				}
				// Verificar si existe el color
				if variant.Color != nil {
					s.products[index].Variants[vIndex].Color.Hex = strings.TrimSpace(variant.Color.Hex)
					s.products[index].Variants[vIndex].Color.Name = strings.TrimSpace(variant.Color.Name)
				}
				// Verificar que el tipo sea valido
				if variant.Type == "" {
					fmt.Printf("El producto %s no tiene tipo de medición\n", product.Name)
				} else {
					variant.Type = strings.ToLower(variant.Type)
					s.products[index].Variants[vIndex].Type = variant.Type
					// Verificar si el tipo esta dentro de los permitidos
					if !slices.Contains(constants.GetValidProductTypes(), variant.Type) {
						fmt.Printf("El producto %s tiene un tipo de medición desconocido (%s)\n", product.Name, variant.Type)
					}
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
