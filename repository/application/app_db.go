package application

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type appRepositoryDB struct {
	db *mongo.Client
}

func NewAppRepositoryDB(db *mongo.Client) appRepositoryDB {
	return appRepositoryDB{db: db}
}

func (a appRepositoryDB) GetAllAppsByUserId(userId string) ([]*Application, error) {
	coll := a.db.Database("Authentacle").Collection("application")
	filter := bson.D{{Key: "userId", Value: userId}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	// var applications []*Application
	for cursor.Next(context.TODO()) {
		var result bson.D
		if err := cursor.Decode(&result); err != nil {

			return nil, err
		}
		// applications = append(applications, &Application{Id: result.Map()["ObjectId"], Name: result.Map()["name"].(string), Logo: result.Map()["logo"].(string), Key: result.Map()["key"].(string)})
	}
	return nil, nil
}

func (a appRepositoryDB) GetAppById(userId string) (*Application, error) {
	a.db.Database("Authentacle").Collection("application")
	return nil, nil
}

func (a appRepositoryDB) AddApp(application *Application) error {
	coll := a.db.Database("Authentacle").Collection("application")
	doc := bson.D{{Key: "name", Value: application.Name}, {Key: "logo", Value: application.Logo}, {Key: "key", Value: application.Key}, {Key: "userId", Value: application.UserId}}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted a single document: %v", result.InsertedID)
	return nil
}
