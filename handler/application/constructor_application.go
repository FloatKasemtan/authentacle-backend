package application

import (
	"github.com/floatkasemtan/authentacle-service/service/application"
)

type applicationHandler struct {
	applicationService application.ApplicationService
}

func NewAppHandler(applicationService application.ApplicationService) applicationHandler {
	return applicationHandler{applicationService: applicationService}
}
