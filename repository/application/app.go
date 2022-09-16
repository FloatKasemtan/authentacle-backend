package application

import "go.mongodb.org/mongo-driver/bson/primitive"

type Application struct {
	Id     primitive.ObjectID `json:"id"`
	Name   string             `json:"name"`
	Logo   string             `json:"logo"`
	Key    string             `json:"key"`
	UserId primitive.ObjectID `json:"userId"`
}

type Repository interface {
	GetAllAppsByUserId(id string) ([]*Application, error)
	GetAppById(id string) (*Application, error)
	AddApp(*Application) error
}
