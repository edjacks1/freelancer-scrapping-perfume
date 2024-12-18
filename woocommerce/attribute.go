package woocommerce

import (
	"encoding/json"
	"fmt"
	"perfume/domain/dto"
	"strings"

	"github.com/sirupsen/logrus"
)

func (s Service) VerifyAttributes(attributes []dto.Attribute) {
	// Registrar en el log.
	s.logger.Debug(fmt.Sprintf("Se van a verificar %d atributos", len(attributes)))
	// Obtener tamaños
	sizeTerms, err := s.GetAttributeTerms(1)
	// Verificar si existe error
	if err != nil {
		s.logger.WithField("err", err).Error("Ocurrio un error al tratar de obtener los tamaños")
		return
	}
	// Obtener colores
	colorTerms, err := s.GetAttributeTerms(2)
	// Verificar si existe error
	if err != nil {
		s.logger.WithField("err", err).Error("Ocurrio un error al tratar de obtener los colores")
		return
	}
	// Obtener marcas
	brandTerms, err := s.GetAttributeTerms(3)
	// Verificar si existe error
	if err != nil {
		s.logger.WithField("err", err).Error("Ocurrio un error al tratar de obtener las marcas")
		return
	}
	// Iterar attributos de productos escaneados
	for _, attribute := range attributes {
		var exist bool
		var index int
		// Dependiendo del tipo verificar
		switch attribute.Type {
		case dto.AttributeSizeType:
			index = 1
			exist = s.CheckIfAttributeExistInTerms(attribute, *sizeTerms)
			break
		case dto.AttributeBrandType:
			index = 3
			exist = s.CheckIfAttributeExistInTerms(attribute, *brandTerms)
			break
		case dto.AttributeColorType:
			index = 2
			exist = s.CheckIfAttributeExistInTerms(attribute, *colorTerms)
			break
		default:
			exist = false
		}
		// Verificar si existe
		if !exist {
			// Itentando crear termino
			s.logger.WithField("attribute", attribute).Debug(fmt.Sprintf("Creando atributo"))
			//Hacer peticion
			response, err := s.rest.Post(
				dto.NewRequestParams{
					Url:     fmt.Sprintf("products/attributes/%d/terms", index),
					Headers: []dto.RequestHeader{s.GetAuthorizationHeader()},
				},
				dto.WCAttributeTerm{
					Name: attribute.Name,
					Detail: dto.WooVariationSwatches{
						PrimaryColor: attribute.Value,
					},
				},
			)
			// Verificar si ocurrio erro
			if err != nil {
				s.logger.
					WithFields(logrus.Fields{
						"error":     err,
						"attribute": attribute,
					}).
					Error(fmt.Sprintf("Ocurrio un error al crear el atributo"))
				// Continuar con el siguiente
				continue
			}
			// Verificar si fue exitoso
			if response.OkStatus {
				s.logger.WithField("body", string(response.Data)).Debug(fmt.Sprintf("Exito al crear atributo"))
			} else {
				s.logger.
					WithFields(logrus.Fields{
						"body":      string(response.Data),
						"attribute": attribute,
					}).
					Error(fmt.Sprintf("No se pudo crear el atributo"))
			}
		}
	}
}

func (s Service) GetAttributeTerms(id int) (*[]dto.WCAttributeTerm, error) {
	page := 1
	allTerms := []dto.WCAttributeTerm{}
	// Clear bucle
	for {
		// Hacer peticion
		response, err := s.rest.Get(dto.NewRequestParams{
			Url:     fmt.Sprintf("products/attributes/%d/terms?per_page=100&page=%d", id, page),
			Headers: []dto.RequestHeader{s.GetAuthorizationHeader()},
		})
		// Validar si fue exitoso
		if err != nil {
			return nil, fmt.Errorf("Error al realizar peticion:", err)
		}
		// Obtener terminos
		var terms []dto.WCAttributeTerm
		// Decodificar
		if err := json.Unmarshal(response.Data, &terms); err != nil {
			return nil, fmt.Errorf("Error al decodificar JSON:", err)
		}
		// Agregar términos a la lista total
		allTerms = append(allTerms, terms...)
		// Detener si no hay más datos
		if len(terms) < 100 {
			break
		}
		// Incrementar pagina
		page++
	}
	// Regresar terminos
	return &allTerms, nil
}

func (s Service) CheckIfAttributeExistInTerms(attribute dto.Attribute, terms []dto.WCAttributeTerm) bool {
	// Verificar el tipo del atributo
	for _, term := range terms {
		if attribute.Type == dto.AttributeColorType {
			if (strings.ToLower(attribute.Name) == strings.ToLower(term.Name)) && (strings.ToLower(attribute.Value) == strings.ToLower(term.Detail.PrimaryColor)) {
				return true
			}
		} else {
			if strings.ToLower(attribute.Name) == strings.ToLower(term.Name) {
				return true
			}
		}
	}
	// Regresar bandera
	return false
}
