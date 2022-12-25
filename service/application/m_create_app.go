package application

import (
	"github.com/floatkasemtan/authentacle-service/repository/application"
	"github.com/floatkasemtan/authentacle-service/type/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
