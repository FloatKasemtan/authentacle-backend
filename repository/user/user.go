package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Email    string             `json:"email"`
	IsVerify bool               `json:"isVerify"`
	Secret   string             `json:"secret"`
}

type Repository interface {
	SignUp(username string, email string, password string) (string, error)
	SignIn(username string) (*User, error)
	GetById(userId string) (*User, error)
	SendVerificationForm(userId string, secret string) error
}
