package service

import (
	"context"
	"perfume/dao"
	"time"

	"github.com/chromedp/chromedp"
)

type Service struct {
	products []dao.Product
}

func InitService() Service {
	return Service{}
}

func (s Service) InitContext() (context.Context, []context.CancelFunc) {
	// Crear un contexto de Chrome
	ctx, chromedp_cancel := chromedp.NewContext(context.Background())
	// Añadir timeout
	ctx, ctx_cancel := context.WithTimeout(ctx, 15*time.Second)
	//Regresar data
	return ctx, []context.CancelFunc{chromedp_cancel, ctx_cancel}
}

// Cancelar contextos
func (s Service) CancelContexts(cancelFns []context.CancelFunc) {
	for _, cancelFn := range cancelFns {
		cancelFn()
	}
}
