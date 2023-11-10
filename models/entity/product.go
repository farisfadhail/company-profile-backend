package entity

import "time"

type Product struct {
	ID                int            `gorm:"primaryKey"`
	ProductCategoryId int            `gorm:"column:product_category_id"`
	Title             string         `gorm:"size:255;column:title;uniqueIndex"`
	Slug              string         `gorm:"size:255;column:slug;uniqueIndex"`
	Material          string         `gorm:"size:255;column:material"`
	Type              string         `gorm:"size:255;column:type"`
	Static            string         `gorm:"size:255;column:static"`
	Dynamic           string         `gorm:"size:255;column:dynamic"`
	Racking           string         `gorm:"size:255;column:racking;default:'N/A'"`
	TokopediaLink     string         `gorm:"column:tokopedia_link;default:null"`
	ShopeeLink        string         `gorm:"column:shopee_link;default:null"`
	LazadaLink        string         `gorm:"column:lazada_link;default:null"`
	ImageGalleries    []ImageGallery `json:"image_galleries" gorm:"foreignKey:product_id;references:id"`
	CreatedAt         time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}