package application

import (
	"errors"
	"github.com/floatkasemtan/authentacle-service/repository/application"
	"github.com/floatkasemtan/authentacle-service/type/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type applicationService struct {
	applicationRepository application.Repository
}

func NewAppService(applicationRepository application.Repository) applicationService {
	return applicationService{applicationRepository: applicationRepository}
}

func (a applicationService) GetApp(id string, userId string) (*ApplicationResponse, error) {
	app, err := a.applicationRepository.GetAppById(id)
	if err != nil {
		return nil, err
	}
	if app.UserId.Hex() != userId {
		return nil, errors.New("Permission denied")
	}

	return &ApplicationResponse{Name: app.Name, Logo: app.Logo, Key: app.Key}, nil
}

func (a applicationService) GetAllApps(id string) ([]*ApplicationsResponse, error) {

	applications, err := a.applicationRepository.GetAllAppsByUserId(id)
	if err != nil {
		return nil, err
	}

	var applicationResponses []*ApplicationsResponse
	for _, app := range applications {
		applicationResponses = append(applicationResponses, &ApplicationsResponse{Name: app.Name, Logo: app.Logo})
	}

	return applicationResponses, nil
}

func (a applicationService) CreateApp(applicationRequest *request.ApplicationRequest, userId string) error {
	primitiveId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	if err := a.applicationRepository.AddApp(&application.Application{Name: applicationRequest.Name, Logo: applicationRequest.Logo, Key: applicationRequest.UnhashKey, UserId: primitiveId}); err != nil {
		return err
	}
	return nil
}
