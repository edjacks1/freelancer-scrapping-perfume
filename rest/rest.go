package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"perfume/domain/dto"
)

func (rest Rest) NewRequest(params dto.NewRequestParams) (*dto.RequestResponse, error) {
	// generar url nueva
	url := fmt.Sprintf("%s/%s", rest.baseUrl, params.Url)
	// Verificar si existen query params
	if params.QueryParams != nil {
		queryParams, err := rest.ConvertStructToURLValues(params.QueryParams)
		// Verificar si no tiene error
		if err != nil {
			url += fmt.Sprintf("/%s", queryParams)
		} else {
			return nil, fmt.Errorf("Ocurrio un error al tratar de generar los query params")
		}
	}
	// Inicializar request
	req, err := http.NewRequest(params.Method, url, params.Data)
	// Verificar si existe error
	if err != nil {
		return nil, fmt.Errorf("Ocurrio un error al tratar de crear la peticion, err: %s", err.Error())
	}
	// Iterar headers para agregar
	for _, header := range params.Headers {
		req.Header.Add(header.Key, header.Value)
	}
	// Realizar la solicitud HTTP
	client := &http.Client{}
	resp, clientErr := client.Do(req)
	// Verificar estatus de peticion
	if clientErr != nil {
		return nil, fmt.Errorf("Error al realizar la solicitud: %s", clientErr.Error())
	}
	// Cerrar el body al final
	defer resp.Body.Close()
	// Leer la respuesta
	body, bodyErr := io.ReadAll(resp.Body)
	// Verificar que este correcto
	if bodyErr != nil {
		return nil, fmt.Errorf("Error al leer la respuesta: %s", bodyErr.Error())
	}
	// Regresar data
	return &dto.RequestResponse{
		Data:     body,
		Request:  *req,
		Response: *resp,
		OkStatus: (resp.StatusCode >= 200 && resp.StatusCode < 300),
	}, nil
}

func (rest Rest) Get(params dto.NewRequestParams) (*dto.RequestResponse, error) {
	// Asignar tipo de metodo
	params.Method = http.MethodGet
	// Realizar peticion
	return rest.NewRequest(params)
}

func (rest Rest) Post(params dto.NewRequestParams, data interface{}) (*dto.RequestResponse, error) {
	// Verificar si existe data
	if data != nil {
		dataJson, err := json.Marshal(data)
		// Verificar si existe error
		if err != nil {
			return nil, fmt.Errorf("Error al convertir la data: %", err.Error())
		} else {
			params.Data = bytes.NewBuffer(dataJson)
		}
	}
	// Asignar tipo de metodo
	params.Method = http.MethodPost
	params.Headers = append(params.Headers, dto.RequestHeader{Key: "Content-Type", Value: "application/json"})
	// Realizar peticion
	return rest.NewRequest(params)
}
