package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u userRepositoryDB) GetById(id string) (*User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := &User{}
	if err := u.coll.FindByID(objectId, result); err != nil {
		return nil, err
	}

	return result, nil
}
