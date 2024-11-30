package model

import (
	"time"

	"gorm.io/gorm"
)

// 订单状态常量
const (
	OrderStatusPending   = iota + 1 // 待支付
	OrderStatusPaid                 // 已支付
	OrderStatusShipped              // 已发货
	OrderStatusCompleted            // 已完成
	OrderStatusCancelled            // 已取消
)

// Order 订单模型
type Order struct {
	ID         uint           `gorm:"primarykey"`          // 订单ID，主键
	UserID     uint           `gorm:"index"`               // 用户ID，外键
	OrderNo    string         `gorm:"uniqueIndex;size:32"` // 订单号，唯一索引
	Status     int            `gorm:"default:1"`           // 订单状态，默认1（待支付）
	TotalPrice float64        `gorm:"type:decimal(10,2)"`  // 订单总价
	Items      []OrderItem    // 订单项，一对多关系
	CreatedAt  time.Time      // 创建时间
	UpdatedAt  time.Time      // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"index"` // 软删除时间
}

// OrderItem 订单项模型
type OrderItem struct {
	ID        uint    `gorm:"primarykey"` // 订单项ID，主键
	OrderID   uint    `gorm:"index"`      // 订单ID，外键
	ProductID uint    `gorm:"index"`      // 商品ID，外键
	Quantity  int     // 购买数量
	Price     float64 `gorm:"type:decimal(10,2)"` // 商品单价
}
