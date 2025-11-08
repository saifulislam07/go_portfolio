package models

import (
	"time"

	"gorm.io/gorm"
)

type Contacts struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);not null"`
	Mobile    string `gorm:"type:varchar(20);not null"`
	Message   string `gorm:"type:text;not null"`
	Status    string `gorm:"type:enum('pending','read','cancelled');default:'pending'"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func ContactsFindLatest() (Contacts, error) {
	var contacts Contacts
	result := DB.Order("id desc").First(&contacts)
	return contacts, result.Error
}

func ContactstList() ([]Contacts, error) {
	var contacts []Contacts
	result := DB.Order("id desc").Find(&contacts)
	return contacts, result.Error
}

func ContactsFind(id uint64) *Contacts {
	var contact Contacts
	result := DB.First(&contact, id)
	if result.Error != nil {
		return nil
	}
	return &contact
}
