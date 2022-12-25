package user

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (u userRepositoryDB) CreateUser(username string, email string, password string, secret string) (string, error) {
	users := &[]User{}
	filter := bson.D{{Key: "username", Value: username}}
	if err := u.coll.SimpleFind(users, filter); err != nil {
		return "", err
	}

	if len(*users) > 0 {
		return "", errors.New("user already exist")
	}

	filter = bson.D{{Key: "email", Value: email}}
	if err := u.coll.SimpleFind(users, filter); err != nil {
		return "", err
	}

	if len(*users) > 0 {
		return "", errors.New("email already exist")
	}

	user := &User{
		Username: username,
		Password: password,
		Email:    email,
		IsVerify: false,
		Secret:   secret,
	}
	if err := u.coll.Create(user); err != nil {
		return "", err
	}

	id := user.ID.Hex()
	return id, nil
}
