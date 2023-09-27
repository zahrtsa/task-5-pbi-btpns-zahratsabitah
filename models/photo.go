package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	// ID       uint   `gorm:"primaryKey" json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `gorm:"type:varchar(255)" json:"photo_url"`
	UserID   uint   `json:"users_id"` // set userID same as ID from model user
	// User 	User 
}