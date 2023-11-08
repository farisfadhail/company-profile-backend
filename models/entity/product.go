package entity

import "time"

type Product struct {
	ID                int            `gorm:"primaryKey"`
	ProductCategoryId int            `gorm:"column:product_category_id"`
	Title             string         `gorm:"size:255;column:title;uniqueIndex"`
	Material          string         `gorm:"size:255;column:material"`
	Type              string         `gorm:"size:255;column:type"`
	Static            string         `gorm:"size:255;column:static"`
	Dynamic           string         `gorm:"size:255;column:dynamic"`
	Racking           string         `gorm:"size:255;column:racking"`
	TokopediaLink     string         `gorm:"column:tokopedia_link"`
	ShopeeLink        string         `gorm:"column:shopee_link"`
	LazadaLink        string         `gorm:"column:lazada_link"`
	ImageGalleries    []ImageGallery `gorm:"foreignKey:product_id;references:id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}