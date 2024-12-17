package dao

type ProductVariantType string

const (
	ProductVariantMlType   ProductVariantType = "ML"
	ProductVariantUnitType ProductVariantType = "UNIT"
)

type ProductVariant struct {
	ID            string               `json:"id"`
	Type          string               `json:"type"`
	Color         *ProductVariantColor `json:"color,omitempty"`
	Price         string               `json:"price"`
	Photos        []string             `json:"photos"`
	IsActive      bool                 `json:"is_active"`
	Quantity      string               `json:"quantity"`
	Original      string               `json:"original"`
	Description   string               `json:"description"`
	DiscountPrice string               `json:"discount_price"`
}
