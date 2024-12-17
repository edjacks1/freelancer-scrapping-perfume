package service

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"perfume/dao"
)

// Crear archivo
func (s Service) CreateFile() {
	// Convertir el arreglo a JSON con formato legible
	data, err := json.MarshalIndent(s.products, "", "  ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}
	// Crear y abrir un archivo para escribir el JSON
	file, err := os.Create("products.json")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close() // Asegurarse de cerrar el archivo
	// Escribir el JSON en el archivo
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}
}

func (s *Service) LoadFile() error {
	// Abrir el archivo JSON
	file, err := os.Open("products.json")
	// Obtener archivo
	if err != nil {
		return fmt.Errorf("error al abrir el archivo: %w", err)
	}
	defer file.Close() // Asegurarse de cerrar el archivo
	// Leer el contenido del archivo
	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error al leer el archivo: %w", err)
	}
	// Deserializar el JSON en el arreglo de productos
	var products []dao.Product // Asegúrate de tener definida tu estructura `Product`
	if err := json.Unmarshal(data, &products); err != nil {
		return fmt.Errorf("error al deserializar el JSON: %w", err)
	}
	// Asignar los productos cargados al campo del servicio
	s.products = products
	// Regresar sin error
	return nil
}

func (s Service) GetLogFolderPath() string {
	return s.logFolderPath
}
