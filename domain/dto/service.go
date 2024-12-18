package dto

import (
	"context"
)

type Service interface {
	// AddProduct(product dao.Product)
	InitContext() (context.Context, []context.CancelFunc)
	CancelContexts(cancelFns []context.CancelFunc)
	ShowTotalProducts()

	RemoveDuplicates(strings []string) []string
}
