package entity

import "time"

type ImageGallery struct {
	ID        int       `gorm:"primaryKey"`
	ProductId int       `gorm:"column:product_id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}