package db

import (
	"github.com/floatkasemtan/authentacle-service/repository/application"
	"github.com/floatkasemtan/authentacle-service/repository/user"
	"github.com/kamva/mgm/v3"
)

var UserColl *mgm.Collection
var ApplicationColl *mgm.Collection

func initCollection() {
	UserColl = mgm.Coll(new(user.User))
	ApplicationColl = mgm.Coll(new(application.Application))
}
