package pages_club

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

func (p Page) GetProductDetail(url string, isActive bool) {
	ctx, cancelFns := p.svc.InitContext()
	// Terminar funciones
	defer p.svc.CancelContexts(cancelFns)
	// Variable para almacenar los datos del producto
	var product dao.Product
	// Ejecutar todas las tareas de chromedp en una sola llamada
	err := chromedp.Run(ctx,
		// Navegar a la página
		chromedp.Navigate(url),
		// Extraer el detalle del producto
		chromedp.WaitVisible(`h1.titleProduct`, chromedp.ByQuery),
		// Obtener imagenes
		chromedp.Evaluate(`(() => {
			// Inicializar producto
			let product = {
				"name": "",
				"brand": "",
				"variants": [],
				"category": "",
			};
			// Verificar si existe listado
			if( Object.keys(productList).length ){
				let isFirstIteration = true;
				// Iterar listado de skus
				for( let key of Object.keys(productList) ){
					// Verificar si es la primera iteracion para poblar valores
					if( isFirstIteration ){
						let categories = connectifData[key].categories[0].split("/");
						// Inicializar producto
						product = {
							...product,
							"name": connectifData[key].name,
							"brand": connectifData[key].brand,
							"category": categories[categories.length - 1],
						}
						// Actualizar bandera
						isFirstIteration = false;
					}
					// Obtener tipo
					let type = productList[key].variant;
					// Inicializar variante
					let variant = {
						"id": connectifData[key].productId,
						"type": "",
						"price": connectifData[key].unitPriceOriginal + "",
						"quantity": "",
						"is_active" : true,
						"photos" : [],
						"description" : connectifData[key].description,
						"discount_price" : connectifData[key].unitPrice + ""
					}
					// Separar datos
					let details = type.split(" ");
					// Verificar el tamaño
					if( details.length >= 2){
						// Verificar el tipo
						switch( details[ details.length - 1 ].toLowerCase() ){
							case "ml": variant.type = "ML"; break;
							case "u" : variant.type = "U"; break;
							default  : variant.type = "U"; break;
						}
						// Actualizar cantidad
						variant.quantity = details[0];
					}
					// Agregar imagenes
					if( dictColores[key] ){
						variant.photos.push(dictColores[key].Grande)
						// Verificar si existen mas imagenes
						if( Object.keys(dictColores[key].Alts).length ){
							variant.photos.push( ...Object.values(dictColores[key].Alts).map( img => img.Grande ) )
						}
					}
					// Adjuntar variante
					product.variants.push(variant)
				}
			}
			// Regresar producto
			return product;
		})()`, &product),
	)
	//Verificar si existe un error
	if err != nil {
		fmt.Printf("Ocurrio un error al tratar de consultar la siguiente url: %s\n", url)
		fmt.Printf("El detalle del error es: %v\n\n", err)
	}
	// Mostrar producto
	p.svc.AddProduct(product)
}

// productList
