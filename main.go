package main

import (
	"base-trade-rest/api"
	"base-trade-rest/core/validation"
	"base-trade-rest/database"
)

var PORT = ":9090"

func main() {
	database.StartDB()

	validation.InitCustomValidation()

	app := api.SetupRouter()
	app.Run(PORT)
}
