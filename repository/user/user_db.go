package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryDB struct {
	db *mongo.Client
}

func NewUserRepositoryDB(db *mongo.Client) userRepositoryDB {
	return userRepositoryDB{db: db}
}

func (u userRepositoryDB) SignUp(username string, email string, password string) (string, error) {
	coll := u.db.Database(("Authentacle")).Collection("user")
	doc := bson.D{{Key: "username", Value: username}, {Key: "email", Value: email}, {Key: "password", Value: password}}
	result, err := coll.InsertOne(context.TODO(), doc)

	if err != nil {
		return "", err
	}
	
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

// Return JWT Token
func (u userRepositoryDB) SignIn(username string, password string) (string, error) {
	return "", nil
}

func (u userRepositoryDB) GetById(string) (*User, error) {
	return nil, nil
}

// Generate QRcode for authenticator apps
func (u userRepositoryDB) SendVerificationForm(string) error {
	return nil
}
