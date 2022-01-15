package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (user *User) Migrate() (err error) {
	return DB.AutoMigrate(&User{})
}
func (user *User) Create() (tx *gorm.DB) {
	return DB.Create(&user)
}
func (user *User) Update() (tx *gorm.DB) {
	return DB.Save(&user)
}
