package entity

import "time"

type ProductCategory struct {
	ID               int            `gorm:"primaryKey"`
	ParentCategoryId int            `gorm:"column:parent_category_id"`
	Name             string         `gorm:"size:255;column:name;uniqueIndex"`
	Slug             string         `gorm:"size:255;column:slug;uniqueIndex"`
	Products         []Product      `json:"products" gorm:"foreignKey:product_category_id;references:id"`
	CreatedAt        time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}