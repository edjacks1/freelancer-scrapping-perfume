package service

import (
	"context"
	"perfume/domain/dao"
	"time"

	"github.com/chromedp/chromedp"
)

type Service struct {
	products      []dao.Product
	wrongUrls     []string
	logFolderPath string
}

func InitService() Service {
	return Service{
		wrongUrls:     []string{},
		logFolderPath: "logs",
	}
}

func (s Service) InitContext(seconds time.Duration) (context.Context, []context.CancelFunc) {
	// Crear un contexto de Chrome
	ctx, chromedp_cancel := chromedp.NewContext(context.Background())
	// Añadir timeout
	ctx, ctx_cancel := context.WithTimeout(ctx, seconds*time.Second)
	//Regresar data
	return ctx, []context.CancelFunc{chromedp_cancel, ctx_cancel}
}

// Cancelar contextos
func (s Service) CancelContexts(cancelFns []context.CancelFunc) {
	for _, cancelFn := range cancelFns {
		cancelFn()
	}
}
