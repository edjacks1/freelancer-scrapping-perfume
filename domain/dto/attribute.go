package dto

type AttributeType string

const (
	AttributeSizeType  AttributeType = "size"
	AttributeColorType AttributeType = "color"
	AttributeBrandType AttributeType = "brand"
)

type Attribute struct {
	Name  string        `json:"name"`
	Type  AttributeType `json:"type"`
	Value string        `json:"value"`
}
