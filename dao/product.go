package dao

type ProductVariantType string

const (
	ProductVariantMlType   ProductVariantType = "ML"
	ProductVariantUnitType ProductVariantType = "UNIT"
)

type Product struct {
	Name     string           `json:"name"`
	Detail   string           `json:"detail"`
	Category string           `json:"category"`
	Variants []ProductVariant `json:"variants"`
}

type ProductVariant struct {
	Type     ProductVariantType `json:"type"`
	Price    string             `json:"price"`
	Quantity string             `json:"quantity"`
	Photos   []string           `json:"photons"`
}
