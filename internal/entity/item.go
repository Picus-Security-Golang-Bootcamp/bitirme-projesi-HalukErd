package entity

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID string `gorm:"primaryKey"`
}

func (u *Item) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewV1().String()
	return nil
}
