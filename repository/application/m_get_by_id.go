package application

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a appRepositoryDB) GetAppById(id string) (*Application, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	applciation := &Application{}
	if err = a.coll.FindByID(objectId, applciation); err != nil {
		return nil, err
	}

	return applciation, nil
}
