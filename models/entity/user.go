package entity

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;column:name"`
	Email     string    `gorm:"size:255;column:email;uniqueIndex"`
	Password  string    `json:"-" gorm:"size:255;column:password"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}