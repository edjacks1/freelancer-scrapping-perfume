package woocommerce

import (
	"encoding/base64"
	"fmt"
	"perfume/domain/dto"
)

func (svc Service) GetAuthorizationHeader() dto.RequestHeader {
	authHeader := base64.StdEncoding.EncodeToString([]byte(svc.client + ":" + svc.secret))
	// Agregar header
	return dto.RequestHeader{
		Key:   "Authorization",
		Value: fmt.Sprintf("Basic %s", authHeader),
	}
}
