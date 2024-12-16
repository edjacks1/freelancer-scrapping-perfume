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
	svc.CalculateDuration(startDate, endDate)
}

func searchInDruniPage() {
	//Obtener pagina
	druniPage := pages_druni.InitPage(&svc)
	//Obtener productos
	druniPage.GetProductDetail("https://www.druni.es/lady-secret-aquarius-cosmetics-eau-toilette-mujer", true)
	druniPage.GetList("https://www.druni.es/marcas/aqc-fragances")
}

func searchInClubPage() {
	fmt.Println("Se ha consultado pagina club")
	// Obtener pagina
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
	})

	// Obtener productos
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/4711/4711-eau-de-colonia/p_89291/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/boss-bottled-eau-de-toilette-vaporizador/p_23010/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/boss-bottled-absolu-eau-de-parfum-vaporizador/p_30005/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/boss-bottled-after-shave/p_23020/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/boss-bottled-desodorante-vaporizador/p_23150/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/boss-bottled-desodorante-stick/p_23140/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/hugo-desodorante-stick/p_24530/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/hugo-eau-de-toilette-vaporizador/p_24480/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/hugo-lote-020045/p_020045/?pid=199840")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/hugo-desodorante-vaporizador/p_24520/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/chloe/chloe-signature-eau-de-perfume-vaporizador/p_12040/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/escada/brisa-cubana-eau-de-toilette-vaporizador/p_93558/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/escada/magnetism-eau-de-perfume-vaporizador/p_17030/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/vanderbilt/vanderbilt-eau-de-toilette-vaporizador/p_43000/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/iroha-nature/exfoliating-toner-pre-soaked-pads/p_364366/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/iroha-nature/nature-mask-dragon-fruit--hyaluronic-acid/p_436100/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/iroha-nature/sos-pimple-patches/p_364211/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/iroha-nature/algodon-mask-face--neck-collagen-antiage/p_31549/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/jil-sander/jil-sander-sun-eau-de-toilette-vaporizador/p_32540/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/jil-sander/jil-sander-sun-men-eau-de-toilette-vaporizador/p_32510/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/marc-jacobs/perfect-elixir-eau-de-parfum-vaporizador/p_30003/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/paloma-picasso/paloma-picasso-eau-de-toilette-vaporizador/p_39270/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/rochas/eau-de-rochas-eau-de-toilette-vaporizador/p_40430/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/rochas/rochas-man-eau-de-toilette-vaporizador/p_40510/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/rochas/eau-de-rochas-gel-de-ducha/p_40500/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/rochas/eau-de-rochas-locion-hidratante-corporal/p_40490/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/rochas/rochas-eau-fraiche-eau-de-toilette/p_40610/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/soap--glory/clean-on-me-creamy-clarifying-shower-gel/p_94915/")
	// clubPage.GetProductDetail("https://www.perfumesclub.com/es/jesus-del-pozo/halloween-eau-de-toilette-vaporizador/p_32310/")
}

func searchInTintinPage() {
	//Obtener pagina
	page := pages_tintin.InitPage(&svc)
	page.GetList("https://www.perfumestintin.com/es/99300-flor-d-ametller")
}

// lugepemo2
// la clave es gerardho98 o gerardo98
