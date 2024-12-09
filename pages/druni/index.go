package pages_druni

import (
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

func (p Page) GetProductDetail(url string) {
	ctx, cancelFns := p.svc.InitContext()
	// Terminar funciones
	defer p.svc.CancelContexts(cancelFns)
	// Variable para almacenar los datos del producto
	var variant dao.ProductVariant
	var product dao.Product
	// Ejecutar todas las tareas de chromedp en una sola llamada
	err := chromedp.Run(ctx,
		// Navegar a la página
		chromedp.Navigate(url),
		// Extraer el nombre del producto
		chromedp.WaitVisible(`.product-info-main .product-title .manufacturer .value`, chromedp.ByQuery),
		chromedp.Text(`.product-info-main .product-title .manufacturer .value`, &product.Name, chromedp.NodeVisible),
		// Extraer el detalle del producto
		chromedp.Text(`.product.attribute .value p`, &product.Detail, chromedp.NodeVisible),
		// Extraer categoria
		chromedp.Text(`.product-info-main .page-title-wrapper.product-simple .page-title`, &product.Category, chromedp.NodeVisible),
		// Extraer contenido
		chromedp.AttributeValue(`.product-info-main  .price-box:nth-child(2) meta[itemprop="price"]`, "content", &variant.Price, nil),
		// Obtener imagenes
		chromedp.Evaluate(`(() => {
			const images = Array.from(document.querySelectorAll('.product-info-main .gallery img'));
			return images.map(img => img.src);
		})()`, &variant.Photos),
	)
	//Verificar si existe un error
	if err != nil {
		fmt.Println(err)
	}
	// Asignar variante
	product.Variants = append(product.Variants, variant)
	// Mostrar producto
	p.svc.AddProduct(product)
}
