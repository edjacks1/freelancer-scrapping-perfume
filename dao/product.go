package dao

type ProductVariantType string

const (
	ProductVariantMlType   ProductVariantType = "ML"
	ProductVariantUnitType ProductVariantType = "UNIT"
)

type Product struct {
	Name     string           `json:"name"`
	Brand    string           `json:"brand"`
	Category string           `json:"category"`
	Variants []ProductVariant `json:"variants"`
}

type ProductVariant struct {
	ID            string             `json:"id"`
	Type          ProductVariantType `json:"type"`
	Price         string             `json:"price"`
	Photos        []string           `json:"photos"`
	IsActive      bool               `json:"is_active"`
	Quantity      string             `json:"quantity"`
	Description   string             `json:"description"`
	DiscountPrice string             `json:"discount_price"`
}
