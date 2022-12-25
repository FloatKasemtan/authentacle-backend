package user

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Email            string `json:"email"`
	IsVerify         bool   `json:"isVerify"`
	Secret           string `json:"secret"`
}

type Repository interface {
	CreateUser(username string, email string, password string, secret string) (string, error)
	GetUserByUsername(username string) (*User, error)
	GetById(userId string) (*User, error)
	Verify(userId string) error
}
