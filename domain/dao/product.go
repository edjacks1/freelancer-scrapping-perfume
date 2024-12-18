package dao

type Product struct {
	Name     string           `json:"name"`
	Brand    string           `json:"brand"`
	Category string           `json:"category"`
	Variants []ProductVariant `json:"variants"`
}

// ck_05242721acc6d004dd5f213a784cc90cb00a4628
// cs_df5bc346cb259a9b6ab647276311ad0d9c235a76
