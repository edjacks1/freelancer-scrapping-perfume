package dto

import (
	"io"
	"net/http"
)

type NewRequestParams struct {
	Url         string
	Data        io.Reader
	Method      string
	Headers     []RequestHeader
	QueryParams interface{}
}

type RequestHeader struct {
	Key   string
	Value string
}

type RequestResponse struct {
	Data     []byte
	Request  http.Request
	Response http.Response
	OkStatus bool
}
