package models

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        string     `gorm:"type:text;primary_key;" json:"id"`
	CreatedAt time.Time  `                              json:"createdAt"`
	UpdatedAt time.Time  `                              json:"updatedAt"`
	DeletedAt *time.Time `                              json:"deletedAt" sql:"index"`
}

// BeforeCreate will set a nanoid rather than numeric ID.
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	//1% collision in 1000 years @ 1000 ids/hr, tweak values depending on use-case
	id, err := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-@~", 12)
	base.ID = id

	return err
}
