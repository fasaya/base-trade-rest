package main

import (
	"base-trade-rest/api"
	"base-trade-rest/database"
	"log"

	"github.com/joho/godotenv"
)

var PORT = ":5050"

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	database.StartDB()

	app := api.SetupRouter()
	app.Run(PORT)
}
