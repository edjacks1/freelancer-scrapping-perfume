package service

import (
	"encoding/json"
	"fmt"
	"os"
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
