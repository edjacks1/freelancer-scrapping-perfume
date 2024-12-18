package dto

type Page interface {
	GetList(url string)
	GetProductDetail(url string)
}
