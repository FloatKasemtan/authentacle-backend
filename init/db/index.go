package db

import (
	"github.com/kamva/mgm/v3"
	"log"

	"github.com/floatkasemtan/authentacle-service/init/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

func Initialize() {
	if err := mgm.SetDefaultConfig(nil, config.C.DB_NAME, options.Client().ApplyURI(config.C.DB_HOST)); err != nil {
		log.Panic("Unable to set MGM configuration: " + err.Error())
	}

	if _, client, database, err := mgm.DefaultConfigs(); err != nil {
		log.Panic("Unable to start MGM: " + err.Error())
	} else {
		Client = client
		Database = database
	}

	initCollection()
}
