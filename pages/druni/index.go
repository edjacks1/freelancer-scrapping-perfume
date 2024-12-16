package pages_druni

import (
	"encoding/json"
	"fmt"
	"perfume/dao"
	"perfume/service"

	"github.com/chromedp/chromedp"
)

type Page struct {
	svc *service.Service
}

func InitPage(svc *service.Service) Page {
	return Page{svc}
}

func (p Page) GetList(url string) {
	ctx, cancelFns := p.svc.InitContext(60)
	// Terminar funciones
	defer p.svc.CancelContexts(cancelFns)
	// Obtener urs
	var links []struct {
		Link       string `json:"link"`
		IsDisabled bool   `json:"isDisabled"`
	}
	//
	err := chromedp.Run(ctx,
		// Navegar a la página
		chromedp.Navigate(url),
		// Extraer el detalle del producto
		chromedp.WaitVisible(`.products.list .product-item-photo`, chromedp.ByQuery),
		chromedp.Evaluate(`(() => {
			return Array.from(document.querySelectorAll('.products-grid .products.list .product-item')).map( el => {
				let link = el.querySelector(".product-item-photo");
				let isDisabled = el.querySelector("form > div:nth-child(2) div a") != null;
				// Mostrar data
				return {
					link       : link.href, 
					isDisabled : isDisabled
				}
			})
		})()`, &links),
	)
	//Verificar si existe un error
	if err != nil {
		fmt.Println(err)
	}
	// Imprimir links
	for _, link := range links {
		p.GetProductDetail(link.Link, !link.IsDisabled)
	}
}

func (p Page) GetProductDetail(url string, isActive bool) {
	ctx, cancelFns := p.svc.InitContext(60)
	// Terminar funciones
	defer p.svc.CancelContexts(cancelFns)
	// Variable para almacenar los datos del producto
	var variant dao.ProductVariant
	var product dao.Product
	var productJSON string
	// Verificar si esta activo
	variant.IsActive = isActive
	// Ejecutar todas las tareas de chromedp en una sola llamada
	err := chromedp.Run(ctx,
		// Navegar a la página
		chromedp.Navigate(url),
		// Extraer el detalle del producto
		chromedp.WaitVisible(`.product.attribute .value p`, chromedp.ByQuery),
		chromedp.Text(`.product.attribute .value p`, &variant.Description, chromedp.NodeVisible),
		// Extraer cantidad
		chromedp.Text(`.product-info-main .page-title-wrapper.product-simple .simple-format`, &variant.Quantity, chromedp.NodeVisible),
		// Extraer contenido
		chromedp.AttributeValue(`.product-info-main  .price-box:nth-child(2) meta[itemprop="price"]`, "content", &variant.Price, nil),
		// Obtener imagenes
		chromedp.Evaluate(`(() => {
			const images = Array.from(document.querySelectorAll('.product-info-main .gallery img'));
			return images.map(img => img.src);
		})()`, &variant.Photos),
		// Obtener el producto en crudo,
		chromedp.Evaluate(`JSON.stringify(productDetail);`, &productJSON),
	)
	//Verificar si existe un error
	if err != nil {
		fmt.Println(err)
	}
	// Verificar si existe json
	if err := json.Unmarshal([]byte(productJSON), &product); err != nil {
		fmt.Println("Error al decodificar JSON: %v", err)
	}
	// Quitar duplicados
	variant.Photos = p.svc.RemoveDuplicates(variant.Photos)
	variant.DiscountPrice = variant.Price
	variant.Quantity, variant.Type = p.svc.GetProductVariantType(variant.Quantity)
	// Asignar variante
	product.Variants = append(product.Variants, variant)
	// Mostrar producto
	p.svc.AddProduct(product)
}
