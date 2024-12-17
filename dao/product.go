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

type ProductVariantColor struct {
	Hex  string `json:"hex"`
	Name string `json:"name"`
}

// ck_05242721acc6d004dd5f213a784cc90cb00a4628
// cs_df5bc346cb259a9b6ab647276311ad0d9c235a76
