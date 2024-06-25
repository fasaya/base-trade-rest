package main

import (
	"base-trade-rest/api"
	"base-trade-rest/database"
)

var PORT = ":5050"

func main() {
	database.StartDB()

	app := api.SetupRouter()
	app.Run(PORT)
}
