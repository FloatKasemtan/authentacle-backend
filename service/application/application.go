package application

import (
	"github.com/floatkasemtan/authentacle-service/type/request"
)

type ApplicationResponse struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
	Key  string `json:"key"`
}

type ApplicationsResponse struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type ApplicationService interface {
	GetAllApps(id string) ([]*ApplicationsResponse, error)
	GetApp(id string, userId string) (*ApplicationResponse, error)
	CreateApp(applicationRequest *request.ApplicationRequest, userId string) error
}
