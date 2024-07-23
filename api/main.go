package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/pvfm/enube/api/database"
	"github.com/pvfm/enube/api/routes"
)

	// @title          Enube Importer and Api
	// @version        1.0.0
	// @description    Enube Importer and Api
	// @termsOfService http://swagger.io/terms/

	// @contact.name   Pedro Monteiro
	// @contact.url    github.com/pvfm
	// @contact.email  pvmonteiro26@gmail.com

	// @host     localhost:9292
	// @basepatha /api/v1

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.StartDB()

	routes.Run()
}
