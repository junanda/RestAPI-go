package database

import (
	"fmt"
	"log"

	"github.com/junanda/golang-aa/config"
	"github.com/junanda/golang-aa/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed Connect to Database Mysql: ", err.Error())
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to run Migration: ", err.Error())
	}

	log.Println("Migrate Database")

	return db
}
