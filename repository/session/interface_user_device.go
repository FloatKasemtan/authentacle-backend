package session

import "github.com/kamva/mgm/v3"

type Session struct {
	mgm.DefaultModel `bson:",inline"`
	UserId           string   `json:"user_id" bson:"user_id"`
	OnlineSession    []string `json:"online_session" bson:"online_session"`
}

type Repository interface {
	CreateSession(userId string, userAgent string) error
	GetUserDeviceById(userDeviceId string) (*Session, error)
	GetUserDeviceByUserId(userId string) ([]Session, error)
	VerifyUserDevice(userDeviceId string) error
	RemoveUserDeviceById(userDeviceId string) error
}

// c.Request.UserAgent()
