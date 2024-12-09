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
	druniPage.GetProductDetail("https://www.druni.es/lady-secret-aquarius-cosmetics-eau-toilette-mujer", true)
	druniPage.GetList("https://www.druni.es/marcas/aqc-fragances")
	//Imprimir el total de productos
	svc.ShowTotalProducts()
	svc.CreateFile()
}

// lugepemo2
// la clave es gerardho98 o gerardo98
