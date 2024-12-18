package dto

type WooVariationSwatches struct {
	PrimaryColor string `json:"primary_color"`
}

// Estructura principal para el atributo
type WCAttributeTerm struct {
	ID     int                  `json:"id,omitempty"`
	Name   string               `json:"name"`
	Slug   string               `json:"slug,omitempty"`
	Detail WooVariationSwatches `json:"woo_variation_swatches"`
}
