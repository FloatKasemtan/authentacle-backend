package main

import (
	"github.com/floatkasemtan/authentacle-service/init/db"
	"github.com/floatkasemtan/authentacle-service/init/gin"
	"github.com/floatkasemtan/authentacle-service/init/redis"
)

func main() {
	db.Initialize()
	redis.Initialize()
	gin.Initialize()
}
