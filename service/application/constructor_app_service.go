package application

import (
	"github.com/floatkasemtan/authentacle-service/repository/application"
)

type applicationService struct {
	applicationRepository application.Repository
}

func NewAppService(applicationRepository application.Repository) applicationService {
	return applicationService{applicationRepository: applicationRepository}
}
