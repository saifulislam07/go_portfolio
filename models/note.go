package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NotesAll() *[]Note {
	var notes []Note
	DB.Where("deleted_at IS NULL").Order("updated_at DESC").Find(&notes)
	return &notes
}

func NotesCreate(name string, content string) *Note {
	entry := Note{Name: name, Content: content}
	DB.Create(&entry)
	return &entry
}

// func NotesFind(id uint64) *Note {
// 	var note Note
// 	DB.Where("id = ?", id).First(&note)
// 	return &note
// }

func NotesFind(id uint64) *Note {
	var note Note
	result := DB.First(&note, id)
	if result.Error != nil {
		return nil
	}
	return &note
}
