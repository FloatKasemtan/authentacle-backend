package user

import (
	"github.com/kamva/mgm/v3"
)

type userRepositoryDB struct {
	coll *mgm.Collection
}

func NewUserRepositoryDB(coll *mgm.Collection) userRepositoryDB {
	return userRepositoryDB{coll: coll}
}
