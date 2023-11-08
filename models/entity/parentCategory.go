package entity

import "time"

type ParentCategory struct {
	ID                int               `gorm:"primaryKey"`
	Name              string            `gorm:"size:255;column:name;uniqueIndex"`
	Slug              string            `gorm:"size:255;column:slug"`
	ProductCategories []ProductCategory `json:"-" gorm:"foreignKey:parent_category_id;references:id"`
	CreatedAt         time.Time         `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt         time.Time         `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}