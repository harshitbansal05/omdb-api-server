package main

import (
	"github.com/harshitbansal05/omdb-api-server/routes"
)

func main() {
	router := routes.SetupRouter()

	router.Run("localhost:8080")
}
