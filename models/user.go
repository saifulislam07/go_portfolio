package models

import (
	"hello_gin/helpers"
	"time"
)

type User struct {
	ID        uint64 `gorm:"primaryKey"`
	Username  string `gorm:"size:64"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func UserAvailability(email string) bool {
	var user User
	DB.Where("username = ?", email).First(&User{})
	return (user.ID == 0)
}

func UserCreate(email string, password string) *User {
	hshPasswd, _ := helpers.HashPassword(password)

	entry := User{
		Username: email,
		Password: hshPasswd,
	}
	DB.Create(&entry)
	return &entry
}

func UserFind(id uint64) *User {
	var user User
	result := DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil
	}
	return &user
}

func UserCheck(email string, password string) *User {
	var user User
	DB.Where("username = ?", email).First(&user)
	if user.ID == 0 {
		return nil
	}

	if !helpers.CheckPasswordHash(password, user.Password) {
		return nil
	}
	return &user
}
