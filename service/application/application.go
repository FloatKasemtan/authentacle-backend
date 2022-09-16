package application

import (
	"github.com/floatkasemtan/authentacle-service/type/request"
)

type ApplicationResponse struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
	Key  string `json:"key"`
}

type ApplicationService interface {
	GetAllApps(id string) ([]*ApplicationResponse, error)
	GetApp(id string) (*ApplicationResponse, error)
	CreateApp(applicationRequest *request.ApplicationRequest, userId string) error
}
