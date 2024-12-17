package main

import (
	"fmt"
	"perfume/service"
	"time"

	pages_club "perfume/pages/club"
	pages_druni "perfume/pages/druni"
	pages_tintin "perfume/pages/tintin"
)

var svc service.Service

func main() {
	// Inicializar servicio
	svc = service.InitService()
	startDate := time.Now()
	//Imprimir el total de productos
	searchInClubPage()
	// searchInDruniPage()
	// searchInTintinPage()
	// Mostrar datos
	svc.ValidateProducts()
	svc.ShowTotalProducts()
	svc.CreateFile()
	// Fecha de fin.
	endDate := time.Now()
	// Mostrar mensaje
	fmt.Println(svc.CalculateDuration(startDate, endDate))
}

func searchInDruniPage() {
	//Obtener pagina
	druniPage := pages_druni.InitPage(&svc)
	//Obtener productos
	druniPage.GetProductDetail("https://www.druni.es/lady-secret-aquarius-cosmetics-eau-toilette-mujer", true)
	druniPage.GetList("https://www.druni.es/marcas/aqc-fragances")
}

func searchInClubPage() {
	clubPage := pages_club.InitPage(&svc)
	// Listado de categorais
	clubPage.MassiveSearch([]string{
		"https://www.perfumesclub.com/es/vanderbilt/m/",
		"https://www.perfumesclub.com/es/jesus-del-pozo/m/",
		"https://www.perfumesclub.com/es/hugo-boss/m/",
		"https://www.perfumesclub.com/es/instituto-espanol/m/",
		"https://www.perfumesclub.com/es/issey-miyake/m/",
		"https://www.perfumesclub.com/es/jean-paul-gaultier/m/",
		"https://www.perfumesclub.com/es/jil-sander/m/",
		"https://www.perfumesclub.com/es/joop/m/",
		"https://www.perfumesclub.com/es/kenzo/m/",
		"https://www.perfumesclub.com/es/marc-jacobs/m/",
		"https://www.perfumesclub.com/es/nivea/m/",
		"https://www.perfumesclub.com/es/origins/m/",
		"https://www.perfumesclub.com/es/skin-generics/m/",
		"https://www.perfumesclub.com/es/axe/m/",
		"https://www.perfumesclub.com/es/cacharel/m/",
		"https://www.perfumesclub.com/es/babaria/m/",
		"https://www.perfumesclub.com/es/essence/m/",
		"https://www.perfumesclub.com/es/colgate/m/",
		"https://www.perfumesclub.com/es/st.-moriz/m/",
		"https://www.perfumesclub.com/es/bondi-sands/m/",
		"https://www.perfumesclub.com/es/aqc-fragrances/m/",
		"https://www.perfumesclub.com/es/4711/m/",
		"https://www.perfumesclub.com/es/chloe/m/",
		"https://www.perfumesclub.com/es/escada/m/",
		"https://www.perfumesclub.com/es/iroha-nature/m/",
		"https://www.perfumesclub.com/es/paloma-picasso/m/",
		"https://www.perfumesclub.com/es/rochas/m/",
		"https://www.perfumesclub.com/es/soap--glory/m/",
	})
}

func searchInTintinPage() {
	//Obtener pagina
	page := pages_tintin.InitPage(&svc)
	page.GetList("https://www.perfumestintin.com/es/99300-flor-d-ametller")
}

// lugepemo2
// la clave es gerardho98 o gerardo98
