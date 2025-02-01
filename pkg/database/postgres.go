package database

import (
	"fmt"
	"log"
	"os"

	"github.com/funukonta/management_toko/internal/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() *gorm.DB {
	db_host := os.Getenv("Db_host")
	db_user := os.Getenv("Db_user")
	db_password := os.Getenv("Db_password")
	db_dbname := os.Getenv("Db_dbname")
	db_port := os.Getenv("Db_port")
	db_sslmode := os.Getenv("Db_sslmode")
	db_timezone := os.Getenv("Db_timezone")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		db_host,
		db_user,
		db_password,
		db_dbname,
		db_port,
		db_sslmode,
		db_timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Println("Error : ", err.Error())
		return nil
	}

	// auto migrate
	db.AutoMigrate(
		models.Users{},
		models.Products{},
	)

	return db

}
