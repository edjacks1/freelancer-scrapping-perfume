package pages_club

import (
	"perfume/logger"
	"perfume/service"

	"github.com/sirupsen/logrus"
)

type Page struct {
	svc        *service.Service
	logger     *logrus.Logger
	totalTries int
}

func InitPage(svc *service.Service) Page {
	return Page{
		svc:        svc,
		logger:     logger.Create(svc.GetLogFolderPath(), "club", ""),
		totalTries: 3,
	}
}
