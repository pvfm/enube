package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/pvfm/enube/api/database/migrations"
)

var db *gorm.DB

func StartDB() {
	hostDB := os.Getenv("DB_HOST")
	userDB := os.Getenv("DB_USER")
	passwordDB := os.Getenv("DB_PASSWORD")
	nameDB := os.Getenv("DB_NAME")
	portDB := os.Getenv("DB_PORT")

	dsn := "host=" + hostDB + " user=" + userDB + " password=" + passwordDB + " dbname=" + nameDB + " port=" + portDB + " sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error to connect to database: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
