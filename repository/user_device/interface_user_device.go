package user_device

import "github.com/kamva/mgm/v3"

type UserDevice struct {
	mgm.DefaultModel `bson:",inline"`
	UserId           string `json:"user_id" bson:"user_id"`
	UserAgent        string `json:"user_agent" bson:"user_agent"`
	Verified         bool   `json:"verified" bson:"verified"`
}

// c.Request.UserAgent()
