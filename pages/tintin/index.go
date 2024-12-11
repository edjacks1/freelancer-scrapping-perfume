package pages_tintin

import (
	"fmt"
	"perfume/dao"
	"perfume/service"
	"strings"

	"github.com/chromedp/chromedp"
)

type Page struct {
	svc *service.Service
}

func InitPage(svc *service.Service) Page {
	return Page{svc}
}

func (p Page) GetProductDetail(url string) {
	status := ""
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
		chromedp.WaitVisible(`.ctd_attribute_data .ctd_product_name .ctd_pname`, chromedp.ByQuery),
		chromedp.Text(`.ctd_attribute_data .ctd_product_name .ctd_pname`, &product.Name, chromedp.NodeVisible),
		// Extraer cantidad
		chromedp.Text(`.ctd_attribute_data .ctd_cname`, &variant.Quantity, chromedp.NodeVisible),
		// Extraer precio real
		chromedp.Text(`.ctd_price .wholesale_price`, &variant.Price, chromedp.NodeVisible),
		// Extraer precio descuento
		chromedp.Text(`.ctd_price .price`, &variant.DiscountPrice, chromedp.NodeVisible),
		// Obtener la marca
		chromedp.Text(`.poduct_details_right_col .product_title`, &product.Brand, chromedp.NodeVisible),
		// Obtener la categoria
		chromedp.Text(`.poduct_details_right_col .product_type`, &product.Category, chromedp.NodeVisible),
		// Obtener descripcion
		chromedp.Text(`#idTab1 h3`, &variant.Description, chromedp.NodeVisible),
		// Obtener status
		chromedp.Text(`.ctd_attribute_data .ctd_combination`, &status, chromedp.NodeVisible),
		// Obtener fotos
		chromedp.Evaluate(`(() => {
			const images = Array.from(document.querySelectorAll('.woocommerce div.product div.images img'));
			return images.map(img => img.src);
		})()`, &variant.Photos),
	)
	//Verificar si existe un error
	if err != nil {
		fmt.Println(err)
	}
	// Quitar duplicados
	variant.Photos = p.svc.RemoveDuplicates(variant.Photos)
	// Verificar si esta activo
	variant.IsActive = strings.Contains(strings.ToLower(status), "disponible")
	// Obtener cantidad y tipo
	variant.Quantity, variant.Type = p.svc.GetProductVariantType(variant.Quantity)
	// Limpiar precios
	variant.Price = p.svc.GetDigitsFromString(variant.Price)
	variant.DiscountPrice = p.svc.GetDigitsFromString(variant.DiscountPrice)
	// Agregar variante al producto
	product.Variants = append(product.Variants, variant)
	// Agregar producto a la lista de productos
	p.svc.AddProduct(product)
}
