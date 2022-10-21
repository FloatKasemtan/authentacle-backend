package user

import (
	"context"
	"errors"

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

func (u userRepositoryDB) SignUp(username string, email string, password string, secret string) (string, error) {
	coll := u.db.Database(("Authentacle")).Collection("user")

	var checkReplica bson.D

	filter := bson.D{{Key: "username", Value: username}}
	coll.FindOne(context.TODO(), filter).Decode(&checkReplica)
	if len(checkReplica.Map()) > 0 {
		return "", errors.New("user already exist")
	}

	filter = bson.D{{Key: "email", Value: email}}
	coll.FindOne(context.TODO(), filter).Decode(&checkReplica)
	if len(checkReplica.Map()) > 0 {
		return "", errors.New("email already exist")
	}

	doc := bson.D{{Key: "username", Value: username}, {Key: "email", Value: email}, {Key: "password", Value: password}, {Key: "isVerify", Value: false}, {Key: "secret", Value: secret}}
	result, err := coll.InsertOne(context.TODO(), doc)

	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

// Return JWT Token
func (u userRepositoryDB) SignIn(username string) (*User, error) {
	coll := u.db.Database(("Authentacle")).Collection("user")
	filter := bson.D{{Key: "username", Value: username}}

	var result bson.D
	if err := coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return nil, errors.New("user not found")
	}

	if !result.Map()["isVerify"].(bool) {
		return nil, errors.New("please verify your account first")
	}

	return &User{
		Id:       result.Map()["_id"].(primitive.ObjectID),
		Username: result.Map()["username"].(string),
		Password: result.Map()["password"].(string),
		Email:    result.Map()["email"].(string),
		Secret:   result.Map()["secret"].(string),
	}, nil
}

func (u userRepositoryDB) GetById(id string) (*User, error) {
	coll := u.db.Database(("Authentacle")).Collection("user")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: objectId}}

	var result bson.D
	if err := coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return nil, err
	}

	return &User{
		Id:       result.Map()["_id"].(primitive.ObjectID),
		Username: result.Map()["username"].(string),
		Password: result.Map()["password"].(string),
		Email:    result.Map()["email"].(string),
		IsVerify: result.Map()["isVerify"].(bool),
	}, nil
}

func (u userRepositoryDB) Verify(id string) error {
	coll := u.db.Database(("Authentacle")).Collection("user")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "isVerify", Value: true}}}}

	var result bson.D
	if err := coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(&result); err != nil {
		return err
	}

	return nil
}
