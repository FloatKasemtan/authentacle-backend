package application

import (
	"github.com/kamva/mgm/v3"
)

type appRepositoryDB struct {
	coll *mgm.Collection
}

func NewAppRepositoryDB(coll *mgm.Collection) appRepositoryDB {
	return appRepositoryDB{coll: coll}
}
