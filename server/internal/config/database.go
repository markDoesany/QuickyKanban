package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open(os.Getenv("PRIVATE_DSN")), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to  connect to the database")
	}

	DB = db
	fmt.Println("Database Connected Successfully")
}
