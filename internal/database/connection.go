package database

import (
	"book-explorer-es/internal/models"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func InitPostgreSQLConnection() {
	// Get configs
	host := viper.Get("POSTGRES_HOST")
	port := viper.Get("POSTGRES_PORT")
	user := viper.Get("POSTGRES_USER")
	password := viper.Get("POSTGRES_PASSWORD")
	dbName := viper.Get("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}

	PostgresDB = db
	log.Info("PostgreSQL connection initialized")
}

func Migrate() {
	if PostgresDB == nil {
		log.Fatal("No PostgreSQL connection found")
		return
	}

	err := PostgresDB.AutoMigrate(
		&models.Car{},
		&models.Feature{},
		&models.CarModel{},
		&models.Producer{},
	)

	if err != nil {
		log.Fatal("Migrate DB failed")
		return
	}

	log.Info("PostgreSQL auto migrate successfully")
}
