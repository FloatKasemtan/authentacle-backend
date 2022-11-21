package main

import (
	"github.com/floatkasemtan/authentacle-service/init/db"
	"github.com/floatkasemtan/authentacle-service/init/fiber"
)

func main() {
	db.Initialize()
	fiber.Initialize()
}
