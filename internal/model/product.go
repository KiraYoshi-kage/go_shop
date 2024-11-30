package model

import (
	"time"

	"gorm.io/gorm"
)

// Product 商品模型
type Product struct {
	ID          uint           `gorm:"primarykey" json:"id" example:"1"`
	Name        string         `gorm:"size:128;index" json:"name" example:"iPhone 15"`
	Description string         `gorm:"type:text" json:"description" example:"最新款iPhone"`
	Price       float64        `gorm:"type:decimal(10,2)" json:"price" example:"6999.00"`
	Stock       int            `gorm:"default:0" json:"stock" example:"100"`
	Status      int            `gorm:"default:1" json:"status" example:"1"` // 1: 上架 2: 下架
	CategoryID  uint           `gorm:"index" json:"category_id"`
	CreatedAt   time.Time      `json:"created_at" example:"2023-12-20T10:00:00Z"`
	UpdatedAt   time.Time      `json:"updated_at" example:"2023-12-20T10:00:00Z"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
