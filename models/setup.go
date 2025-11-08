package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "gin_notes:password@tcp(127.0.0.1:3306)/gin_notes?parseTime=true"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Failed to connect to database: " + err.Error())
	}

	log.Println("✅ GORM database connection successful!")
}

func DBMigrate() {
	DB.AutoMigrate(&Note{})
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&About{})
	DB.AutoMigrate(&Skils{})
	DB.AutoMigrate(&Interests{})
	DB.AutoMigrate(&Portfolios{})
	DB.AutoMigrate(&Contacts{})
	DB.AutoMigrate(&SocialMedias{})
}
