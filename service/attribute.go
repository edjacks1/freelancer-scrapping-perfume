package service

import (
	"fmt"
	"perfume/domain/dto"
)

func (s Service) GetAllAttributes() []dto.Attribute {
	// Variables
	attributes := []dto.Attribute{}
	// Iterar los productos
	for _, product := range s.products {
		// Agregar marca
		attributes = append(attributes, dto.Attribute{
			Type: dto.AttributeBrandType,
			Name: product.Brand,
		})
		// Iterar las variantes
		for _, variant := range product.Variants {
			// Agregar data
			attributes = append(attributes, dto.Attribute{
				Type: dto.AttributeSizeType,
				Name: variant.Quantity + variant.Type,
			})
			// Verificar si tiene color
			if variant.Color != nil {
				attributes = append(attributes, dto.Attribute{
					Type:  dto.AttributeColorType,
					Name:  variant.Color.Name,
					Value: variant.Color.Hex,
				})
			}
		}
	}
	// Regresar atributos
	return attributes
}

func (s Service) GetUniqueAttributes() []dto.Attribute {
	// Variable donde se almacenaran los valores unicos.
	keys := map[string]string{}
	filteredAttributes := []dto.Attribute{}
	// Iterar atributos
	for _, attribute := range s.GetAllAttributes() {
		var key string
		// Verificar el tipo
		if attribute.Type == dto.AttributeColorType {
			key = fmt.Sprintf("%s - %s", attribute.Name, attribute.Value)
		} else {
			key = attribute.Name
		}
		// Verificar si existe la llave
		if _, ok := keys[key]; !ok {
			keys[key] = "0"
			filteredAttributes = append(filteredAttributes, attribute)
		}
	}
	// Regresar atributos filtrados
	return filteredAttributes
}
