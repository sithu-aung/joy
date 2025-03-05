package models

import (
	"fmt"
	"log"

	"github.com/sithu-aung/joy/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB(config *config.Config) (*gorm.DB, error){
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
           config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		return nil, err
	}

	// Auto Migrate the schema
	// err = db.AutoMigrate(&User{} , &Course{}, &Lesson{})
	// if err != nil {
	// 	return nil, err
	// }

	log.Println("Database connected successfully")
	return db,nil
}