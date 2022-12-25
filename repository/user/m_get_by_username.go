package user

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

// Return JWT Token
func (u userRepositoryDB) GetUserByUsername(username string) (*User, error) {
	filter := bson.D{{Key: "username", Value: username}}
	users := &[]User{}
	if err := u.coll.SimpleFind(users, filter); err != nil {
		return nil, err
	}
	if len(*users) == 0 {
		return nil, errors.New("user not found")
	}
	user := (*users)[0]

	return &user, nil
}
