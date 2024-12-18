package woocommerce

import (
	"perfume/logger"
	"perfume/rest"

	"github.com/sirupsen/logrus"
)

type Service struct {
	rest   rest.Rest
	secret string
	logger *logrus.Logger
	client string
}

func InitService() Service {
	return Service{
		rest:   rest.InitRest("https://perfumsalomar.com/wp-json/wc/v3"),
		logger: logger.Create("logs", "woocommerce", ""),
		secret: "cs_df5bc346cb259a9b6ab647276311ad0d9c235a76",
		client: "ck_05242721acc6d004dd5f213a784cc90cb00a4628",
	}
}
