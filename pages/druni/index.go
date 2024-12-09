package pages_druni

import (
	"encoding/json"
	"fmt"
	"perfume/dao"
	"perfume/service"
	"strings"

	"github.com/chromedp/chromedp"
)

type ProductDetail struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Brand    string `json:"brand"`
	Price    string `json:"price"`
	Variant  string `json:"variant"`
}

type Page struct {
	svc *service.Service
}

func InitPage(svc *service.Service) Page {
	return Page{svc}
}

func (p Page) GetList(url string) {
	ctx, cancelFns := p.svc.InitContext()
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
	ctx, cancelFns := p.svc.InitContext()
	// Terminar funciones
	defer p.svc.CancelContexts(cancelFns)
	// Variable para almacenar los datos del producto
	var variant dao.ProductVariant
	var product dao.Product
	var productJSON string
	// Verificar si esta activo
	product.IsActive = isActive
	// Ejecutar todas las tareas de chromedp en una sola llamada
	err := chromedp.Run(ctx,
		// Navegar a la página
		chromedp.Navigate(url),
		// Extraer el detalle del producto
		chromedp.WaitVisible(`.product.attribute .value p`, chromedp.ByQuery),
		chromedp.Text(`.product.attribute .value p`, &product.Detail, chromedp.NodeVisible),
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
	// Limpiar variante
	variant.Quantity, variant.Type = p.GetVariantQuantityType(variant.Quantity)
	// Asignar variante
	product.Variants = append(product.Variants, variant)
	// Mostrar producto
	p.svc.AddProduct(product)
}

func (p Page) GetVariantQuantityType(quantity string) (string, dao.ProductVariantType) {
	// Remplazar caracteres
	quantity = strings.ReplaceAll(quantity, "| ", "")
	// Verificar si hay tamaño
	if len(quantity) > 0 {
		// Verificar si es unidad
		if strings.ToLower(quantity)[len(quantity)-1] == 'u' {
			return quantity[:len(quantity)-1], dao.ProductVariantUnitType
		}
		// Verificar si el tamaño es mayor a dos
		if len(quantity) > 2 {
			runes := []rune(strings.ToLower(quantity))
			// Verificar si es mililitro
			if string(runes[len(runes)-2:]) == "ml" {
				return quantity[:len(quantity)-2], dao.ProductVariantMlType
			}
		}
	}
	// Return
	return "1", dao.ProductVariantUnitType
}
