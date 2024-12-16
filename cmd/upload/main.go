package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	woocommerceAPIURL = "https://perfumsalomar.com/wp-json/wc/v3/products" // Cambia con la URL de tu tienda
	consumerKey       = "ck_05242721acc6d004dd5f213a784cc90cb00a4628"      // Reemplaza con tu consumer_key
	consumerSecret    = "cs_df5bc346cb259a9b6ab647276311ad0d9c235a76"      // Reemplaza con tu consumer_secret
)

type Product struct {
	Name       string     `json:"name"`
	Brand      string     `json:"brand"`
	Categories []Category `json:"categories"`
	Variants   []Variant  `json:"variants"`
}

type Category struct {
	Name string `json:"name"`
}

type Variant struct {
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	Price       string  `json:"regular_price"`
	SalePrice   string  `json:"sale_price"`
	Quantity    string  `json:"stock_quantity"`
	Description string  `json:"description"`
	Images      []Image `json:"images"`
}

type Image struct {
	Src string `json:"src"`
}

func main() {
	// Datos del producto
	product := Product{
		Name:  "INSTANT MATT make-up setting spray",
		Brand: "Essence",
		Categories: []Category{
			{Name: "Fijadores de Maquillaje"}, // Nombre de la categoría
		},
		Variants: []Variant{
			{
				ID:          "50019106",
				Type:        "ML",
				Price:       "3.79", // Precio normal
				SalePrice:   "3.22", // Precio con descuento
				Quantity:    "50",   // Cantidad en stock
				Description: "Spray Fijador de Maquillaje Instant Matt Make-Up Settingr 50 ML",
				Images: []Image{
					{Src: "https://i1.perfumesclub.com/grande/50019106.jpg"}, // Imagen del producto
				},
			},
		},
	}

	// Codificar los datos a formato JSON
	data, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}

	// Crear la solicitud HTTP
	req, err := http.NewRequest("POST", woocommerceAPIURL, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	// Configurar la autenticación básica usando claves API
	req.SetBasicAuth(consumerKey, consumerSecret)

	// Establecer cabeceras
	req.Header.Set("Content-Type", "application/json")

	// Hacer la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Verificar el estado de la respuesta
	if resp.StatusCode == 201 {
		fmt.Println("Producto creado exitosamente")
	} else {
		fmt.Printf("Error al crear producto: %s\n", resp.Status)
	}
}
