package handler

import "fmt"

type Handler struct {
	url     string
	message string
}

func Create(url, message string) Handler {
	return Handler{
		url:     url,
		message: message,
	}
}

func (handler Handler) Error() string {
	return handler.message + " " + handler.url
}

func (handler Handler) PrintDetails() {
	fmt.Printf("Ocurrio un error al tratar de consultar la siguiente url: %s\n", handler.url)
	fmt.Printf("El detalle del error es: %v\n\n", handler.message)
}
