package db

import (
	"context"
	"log"

	"github.com/floatkasemtan/authentacle-service/init/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func Initialize() {
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.C.DB_HOST))
	if err != nil {
		log.Panic(err.Error())
	}
	DB = db
}
