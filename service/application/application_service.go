package application

import (
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

func (a applicationService) GetApp(id string) (*ApplicationResponse, error) {
	application, err := a.applicationRepository.GetAppById(id)
	if err != nil {
		return nil, err
	}

	return &ApplicationResponse{Name: application.Name, Logo: application.Logo, Key: application.Key}, nil
}

func (a applicationService) GetAllApps(id string) ([]*ApplicationResponse, error) {

	applications, err := a.applicationRepository.GetAllAppsByUserId("")
	if err != nil {
		return nil, err
	}

	var applicationResponses []*ApplicationResponse
	for _, application := range applications {
		applicationResponses = append(applicationResponses, &ApplicationResponse{Name: application.Name, Logo: application.Logo, Key: application.Key})
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
