package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"perfume/dao"
	"time"

	"github.com/chromedp/chromedp"
)

type Service struct {
	products []dao.Product
}

func InitService() Service {
	return Service{}
}

func (s Service) InitContext() (context.Context, []context.CancelFunc) {
	// Crear un contexto de Chrome
	ctx, chromedp_cancel := chromedp.NewContext(context.Background())
	// Añadir timeout
	ctx, ctx_cancel := context.WithTimeout(ctx, 15*time.Second)
	//Regresar data
	return ctx, []context.CancelFunc{chromedp_cancel, ctx_cancel}
}

// Cancelar contextos
func (s Service) CancelContexts(cancelFns []context.CancelFunc) {
	for _, cancelFn := range cancelFns {
		cancelFn()
	}
}

// Añadir producto
func (s *Service) AddProduct(product dao.Product) {
	s.products = append(s.products, product)
}

// Imprimir el total de productos
func (s Service) ShowTotalProducts() {
	fmt.Printf("Total de productos: %d\n", len(s.products))
}

// Crear archivo
func (s Service) CreateFile() {
	// Convertir el arreglo a JSON con formato legible
	data, err := json.MarshalIndent(s.products, "", "  ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}
	// Crear y abrir un archivo para escribir el JSON
	file, err := os.Create("personas.json")
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
