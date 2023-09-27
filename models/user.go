package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username         string    `gorm:"type:varchar(255);" json:"username"`
	Email            string    `gorm:"uniqueIndex;" json:"email"`
	Password         string    `gorm:"type:varchar(255)" binding:"min=6" json:"password"`
	Photo            Photo     `gorm:"constraint:OnDelete:CASCADE;foreignkey:UserID;references:id" json:"photo"` // one to one relationship
	// CreatedAt        time.Time `gorm:"not null"`
	// UpdatedAt        time.Time `gorm:"not null"`
}
