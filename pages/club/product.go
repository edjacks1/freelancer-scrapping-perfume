package pages_club

import (
	"perfume/dao"

	"github.com/chromedp/chromedp"
	"github.com/sirupsen/logrus"
)

func (p Page) GetProductDetail(url string, tries int) {
	// Mostrar log
	p.logger.WithFields(logrus.Fields{"url": url}).Debug("Iniciando scraping del producto.")
	// Obtener contexto
	ctx, cancelFns := p.svc.InitContext(30)
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
            // Funcion para obtener el color 
            rgbToHex = (rgb) => {
                const rgbValues = rgb.match(/\d+/g); // Extrae los valores R, G y B
                // Validar longitud
                if (rgbValues.length < 3){
                    return ""
                }
                // Regresar data
                return "#" + rgbValues
                    .slice(0, 3) // Solo usamos R, G y B, ignorando alpha si existe
                    .map((num) => {
                        const hex = parseInt(num, 10).toString(16); // Convierte cada valor a hexadecimal
                        return hex.length === 1 ? '0' + hex : hex;  // Asegura que tenga 2 caracteres
                    })
                    .join('');
            }
            // Verificar si existe listado
            if (Object.keys(productList).length) {
                let isFirstIteration = true;
                // Iterar listado de skus
                for (let key of Object.keys(productList)) {
                    // Verificar si es la primera iteracion para poblar valores
                    if (isFirstIteration) {
                        // Inicializar producto
                        product = {
                            ...product,
                            "name": connectifData[key].name,
                            "brand": connectifData[key].brand,
                            "category": connectifData[key].categories[0],
                        }
                        // Actualizar bandera
                        isFirstIteration = false;
                    }
                    // Inicializar variante
                    let variant = {
                        "id": connectifData[key].productId,
                        "type": "",
                        "price": connectifData[key].unitPriceOriginal + "",
                        "quantity": "",
                        "original": productList[key].variant,
                        "is_active": true,
                        "photos": [],
                        "description": connectifData[key].description,
                        "discount_price": connectifData[key].unitPrice + ""
                    }
                    // Separar datos
                    let details = variant.original.split(" ");
                    // Verificar el tamaño
                    if (details.length >= 2) {
                        let detailType = details[details.length - 1].toLowerCase();
                        // Verificar el tipo
                        switch (detailType) {
                            case "ml": variant.type = "ML"; break;
                            case "u": variant.type = "U"; break;
                            default: variant.type = detailType; break;
                        }
                        // Actualizar cantidad
                        variant.quantity = details[0];
                    }
                    // Agregar imagenes
                    if (dictColores[key]) {
                        variant.photos.push(dictColores[key].Grande)
                        // Verificar si existen mas imagenes
                        if (Object.keys(dictColores[key].Alts).length) {
                            variant.photos.push(...Object.values(dictColores[key].Alts).map(img => img.Grande))
                        }
                    }
                    // Verificar si existe color
                    let colorElement = document.querySelector('.colorsContainer div[data-id="' + key + '"]')
                    if( colorElement ){
                        // Crear color
                        variant.color = {
                            hex : rgbToHex(window.getComputedStyle(colorElement.querySelector("span")).backgroundColor),
                            name : colorElement.querySelector("em").textContent
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
		fields := logrus.Fields{
			"url":   url,
			"error": err,
			"tries": tries,
		}
		// Verificar la cantidad de intentos
		if tries < p.totalTries {
			// Registrar en el log
			p.logger.WithFields(fields).Debug("Ocurrio un error al tratar de obtener el producto, volviendo a intentar.")
			// Volver a intentar
			p.GetProductDetail(url, tries+1)
			// Finalizar funcion
			return
		} else {
			p.svc.AddWrongUrls(url)
			p.logger.WithFields(fields).Error("Ocurrio un error al tratar de obtener el producto.")
		}
	} else {
		// Mostrar producto
		p.svc.AddProduct(product)
	}
}
