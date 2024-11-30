package repository

import (
	"myshop/internal/model"

	"gorm.io/gorm"
)

// OrderRepository 订单数据访问层
type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository 创建订单仓储实例
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// Create 创建新订单
func (r *OrderRepository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}

// GetByID 根据ID获取订单
func (r *OrderRepository) GetByID(id uint) (*model.Order, error) {
	var order model.Order
	err := r.db.Preload("Items").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// GetByUserID 获取用户的订单列表
func (r *OrderRepository) GetByUserID(userID uint, page, pageSize int) ([]model.Order, int64, error) {
	var orders []model.Order
	var total int64

	if err := r.db.Model(&model.Order{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := r.db.Where("user_id = ?", userID).
		Preload("Items").
		Offset(offset).
		Limit(pageSize).
		Find(&orders).Error

	if err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

// UpdateStatus 更新订单状态
func (r *OrderRepository) UpdateStatus(id uint, status int) error {
	return r.db.Model(&model.Order{}).Where("id = ?", id).Update("status", status).Error
}

// GetDB 获取数据库连接
func (r *OrderRepository) GetDB() *gorm.DB {
	return r.db
}
