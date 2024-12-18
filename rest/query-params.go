package rest

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func (r Rest) ConvertStructToURLValues(data interface{}) (url.Values, error) {
	values := url.Values{}
	v := reflect.ValueOf(data)
	// Verificar si el valor es un struct
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("data debe ser una estructura")
	}
	// Iterar campos
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		// Obtener la etiqueta `url`
		tag := field.Tag.Get("url")
		if tag == "" || tag == "-" {
			continue // Ignorar si no tiene etiqueta `url` o si está marcada para omitir
		}
		// Manejar distintos tipos de datos
		switch value.Kind() {
		case reflect.String:
			values.Add(tag, value.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			values.Add(tag, strconv.FormatInt(value.Int(), 10))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			values.Add(tag, strconv.FormatUint(value.Uint(), 10))
		case reflect.Float32, reflect.Float64:
			values.Add(tag, strconv.FormatFloat(value.Float(), 'f', -1, 64))
		case reflect.Bool:
			values.Add(tag, strconv.FormatBool(value.Bool()))
		default:
			return nil, fmt.Errorf("tipo no soportado: %s", value.Kind())
		}
	}
	// Regresar valores
	return values, nil
}
