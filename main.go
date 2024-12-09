package main

import (
	pages_druni "perfume/pages/druni"
	"perfume/service"
)

func main() {
	svc := service.InitService()
	//Obtener pagina
	druniPage := pages_druni.InitPage(&svc)
	//Obtener productos
	druniPage.GetProductDetail("https://www.druni.es/lady-secret-aquarius-cosmetics-eau-toilette-mujer")
	//Imprimir el total de productos
	svc.ShowTotalProducts()
	svc.CreateFile()
}
