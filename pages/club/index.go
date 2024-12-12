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

func (p Page) GetList(url string) {
	ctx, cancelFns := p.svc.InitContext()
	// Terminar funciones
	defer p.svc.CancelContexts(cancelFns)
	// Obtener urs
	var links []string
	var dummy interface{}
	// Hacer peticiones
	err := chromedp.Run(ctx,
		// Navegar a la página
		chromedp.Navigate(url),
		// Extraer el detalle del producto
		chromedp.WaitVisible(`#listadoProductos .productList a`, chromedp.ByQuery),
		// Cargar todo
		chromedp.Evaluate(`
		    async function refreshAllItems(){
                // Verificar si existe el botón de cargar más
                while( window.getComputedStyle(document.getElementById("btnNewPage")).display != "none" ){
                    // Obtener articulos
                    var $capaArticulos = $("#listadoProductos");
                    // Sumar pagina
                    marcaArticulosPageModel.Pagina += 1;
                    // Hacer consulta
                    try {
                        // Usamos fetch para hacer la solicitud POST
                        const response = await fetch("/es/marca/marcapagina/", {
                            method: "POST",
                            headers: {
                                "Content-Type": "application/json",  // Ajusta según el tipo de contenido que estés enviando
                            },
                            body: JSON.stringify(marcaArticulosPageModel)  // Convierte el objeto a JSON
                        });
                        // Verificar respuesta
                        if (!response.ok) {
                            throw new Error('Error en la respuesta del servidor');
                        }
                        // Obtén los datos de la respuesta
                        const data = await response.text();  // O response.json() si esperas un JSON
                        // Añadir los datos al elemento
                        $capaArticulos.append(data);
                    } catch (error) {
                        console.error("Error al obtener los datos:", error);
                    }
                    // Esperar 2 segundos
                    await new Promise(r => setTimeout(r, 2000));
                } 
                // Agregar elemento nuevo
                const div = document.createElement('div');
                div.id = 'seHaCompletado';  // Asignar el ID
                // Agregar contenido al div
                div.innerHTML = '¡La operación se ha completado!';
                // Añadir el div al body
                document.body.appendChild(div);
            }
            // Recargar items
            refreshAllItems();
		`, &dummy),
		// Esperar 5 segundos
		chromedp.WaitVisible(`#seHaCompletado`, chromedp.ByID),
		// Obtener datos
		chromedp.Evaluate(`(() => {
			return Array.from(document.querySelectorAll('#listadoProductos .productList')).map( el => {
				let link = el.querySelector("a");
				// Mostrar data
				return link.href;
			})
		})()`, &links),
	)
	//Verificar si existe un error
	if err != nil {
		fmt.Println(err)
	}
	// Imprimir links
	for _, link := range links {
		p.GetProductDetail(link)
	}
}

func (p Page) GetProductDetail(url string) {
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
	} else {
		// Mostrar producto
		p.svc.AddProduct(product)
	}
}

// productList
