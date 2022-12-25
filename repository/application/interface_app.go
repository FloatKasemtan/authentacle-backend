package application

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Application struct {
	mgm.DefaultModel `bson:"base,inline"`
	Name             string             `json:"name"`
	Logo             string             `json:"logo"`
	Key              string             `json:"key"`
	UserId           primitive.ObjectID `json:"userId"`
}

type Repository interface {
	GetAllAppsByUserId(id string) ([]*Application, error)
	GetAppById(id string) (*Application, error)
	AddApp(*Application) error
}
