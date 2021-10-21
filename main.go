package main

import (
	"github.com/LeoGardini/simple-go-api/db"
	"github.com/LeoGardini/simple-go-api/db/migrations"
	"github.com/LeoGardini/simple-go-api/server"
)

func main() {
	db.StartDB()
	migrations.RunMigrations(db.GetDatabase())

	server := server.NewServer()
	server.Run()
}
