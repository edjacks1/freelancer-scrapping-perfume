package main

import (
	pages_club "perfume/pages/club"
	pages_druni "perfume/pages/druni"
	pages_tintin "perfume/pages/tintin"
	"perfume/service"
)

var svc service.Service

func main() {
	// Inicializar servicio
	svc = service.InitService()
	//Imprimir el total de productos
	searchInTintinPage()
	// Mostrar datos
	svc.ShowTotalProducts()
	svc.CreateFile()
}

func searchInDruniPage() {
	//Obtener pagina
	druniPage := pages_druni.InitPage(&svc)
	// //Obtener productos
	druniPage.GetProductDetail("https://www.druni.es/lady-secret-aquarius-cosmetics-eau-toilette-mujer", true)
	druniPage.GetList("https://www.druni.es/marcas/aqc-fragances")
}

func searchInClubPage() {
	// Obtener pagina
	clubPage := pages_club.InitPage(&svc)
	// Obtener productos
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/4711/4711-eau-de-colonia/p_89291/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/boss-bottled-eau-de-toilette-vaporizador/p_23010/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/boss-bottled-absolu-eau-de-parfum-vaporizador/p_30005/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/boss-bottled-after-shave/p_23020/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/boss-bottled-desodorante-vaporizador/p_23150/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/boss-bottled-desodorante-stick/p_23140/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/hugo-desodorante-stick/p_24530/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/hugo-eau-de-toilette-vaporizador/p_24480/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/hugo-lote-020045/p_020045/?pid=199840", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/hugo-boss/hugo-desodorante-vaporizador/p_24520/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/chloe/chloe-signature-eau-de-perfume-vaporizador/p_12040/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/escada/brisa-cubana-eau-de-toilette-vaporizador/p_93558/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/escada/magnetism-eau-de-perfume-vaporizador/p_17030/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/vanderbilt/vanderbilt-eau-de-toilette-vaporizador/p_43000/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/iroha-nature/exfoliating-toner-pre-soaked-pads/p_364366/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/iroha-nature/nature-mask-dragon-fruit--hyaluronic-acid/p_436100/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/iroha-nature/sos-pimple-patches/p_364211/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/iroha-nature/algodon-mask-face--neck-collagen-antiage/p_31549/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/jil-sander/jil-sander-sun-eau-de-toilette-vaporizador/p_32540/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/jil-sander/jil-sander-sun-men-eau-de-toilette-vaporizador/p_32510/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/marc-jacobs/perfect-elixir-eau-de-parfum-vaporizador/p_30003/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/paloma-picasso/paloma-picasso-eau-de-toilette-vaporizador/p_39270/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/rochas/eau-de-rochas-eau-de-toilette-vaporizador/p_40430/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/rochas/rochas-man-eau-de-toilette-vaporizador/p_40510/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/rochas/eau-de-rochas-gel-de-ducha/p_40500/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/rochas/eau-de-rochas-locion-hidratante-corporal/p_40490/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/rochas/rochas-eau-fraiche-eau-de-toilette/p_40610/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/soap--glory/clean-on-me-creamy-clarifying-shower-gel/p_94915/", true)
	clubPage.GetProductDetail("https://www.perfumesclub.com/es/jesus-del-pozo/halloween-eau-de-toilette-vaporizador/p_32310/", true)
}

func searchInTintinPage() {
	//Obtener pagina
	page := pages_tintin.InitPage(&svc)
	// //Obtener productos
	page.GetProductDetail("https://www.perfumestintin.com/es/flor-d-ametller/30010001-flor-d-ametller-flor-d-ametler-eau-de-toilette")
	// page.GetProductDetail("https://www.perfumestintin.com/es/99300-flor-d-ametller", true)

}

// lugepemo2
// la clave es gerardho98 o gerardo98
