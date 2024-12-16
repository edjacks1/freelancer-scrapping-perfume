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

// ck_05242721acc6d004dd5f213a784cc90cb00a4628
// cs_df5bc346cb259a9b6ab647276311ad0d9c235a76
