package main

import (
	"github.com/floatkasemtan/authentacle-service/init/db"
	"github.com/floatkasemtan/authentacle-service/init/gin"
)

func main() {
	db.Initialize()
	gin.Initialize()
}
