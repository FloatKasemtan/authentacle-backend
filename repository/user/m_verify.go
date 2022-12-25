package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u userRepositoryDB) Verify(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result := &User{}
	if err := u.coll.FindByID(objectId, result); err != nil {
		return err
	}
	result.IsVerify = true

	if err := u.coll.Update(result); err != nil {
		return err
	}

	return nil
}
