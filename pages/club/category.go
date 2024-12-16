package pages_club

import (
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/sirupsen/logrus"
)

func (p Page) MassiveSearch(urls []string) {
	// Iterar string
	for _, url := range urls {
		p.GetList(url, 0)
	}
}

func (p Page) GetList(url string, tries int) {
	// Si es el primer intento esperar 5 segundos
	if tries == 0 {
		time.Sleep(5 * time.Second)
	}
	// Mostrar log
	p.logger.WithFields(logrus.Fields{"url": url}).Debug("Iniciando scraping de la categoria.")
	// Incializar
	ctx, cancelFns := p.svc.InitContext(120)
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
                let element = document.getElementById("btnNewPage")
                // Verificar si existe el botón de cargar más
                while( element && window.getComputedStyle(element).display != "none" ){
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
		fields := logrus.Fields{
			"url":   url,
			"error": err,
			"tries": tries,
		}
		// Verificar la cantidad de intentos
		if tries < p.totalTries {
			// Registrar en el log
			p.logger.WithFields(fields).Debug("Ocurrio un error al tratar de obtener el listado de la categoria, volviendo a intentar.")
			// Volver a intentar
			p.GetList(url, tries+1)
			// Finalizar funcion
			return
		} else {
			p.logger.WithFields(fields).Error("Ocurrio un error al tratar de obtener el listado de la categoria.")
		}
	}
	// Imprimir el total de productos a buscar
	p.logger.Debug(fmt.Sprintf("Intentando obtener un total de %d productos", len(links)))
	// Imprimir links
	for _, link := range links {
		p.GetProductDetail(link, 0)
	}
}
