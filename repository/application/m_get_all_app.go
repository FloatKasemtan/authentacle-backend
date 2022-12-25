package application

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a appRepositoryDB) GetAllAppsByUserId(userId string) ([]*Application, error) {
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"userId", objectId}}

	applications := &[]*Application{}
	if err := a.coll.SimpleFind(applications, filter); err != nil {
		return nil, err
	}

	return *applications, nil
}
