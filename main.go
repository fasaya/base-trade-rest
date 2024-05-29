package main

import (
	"base-trade-rest/database"
	"base-trade-rest/router"
)

var PORT = ":9090"

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}
