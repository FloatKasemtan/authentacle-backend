package user

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"time"

	"github.com/floatkasemtan/authentacle-service/init/config"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryDB struct {
	db *mongo.Client
}

func NewUserRepositoryDB(db *mongo.Client) userRepositoryDB {
	return userRepositoryDB{db: db}
}

// Return JWT Token
func (u userRepositoryDB) SignUp(username string, email string, password string) (string, error) {
	coll := u.db.Database(("Authentacle")).Collection("user")
	doc := bson.D{{Key: "username", Value: username}, {Key: "email", Value: email}, {Key: "password", Value: password}}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return "", err
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":  result.InsertedID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	spew.Dump(claims)

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.C.JWT_SECRET))

	return t, nil
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
