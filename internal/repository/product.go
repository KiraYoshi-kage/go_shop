package repository

import (
	"myshop/internal/model"

	"gorm.io/gorm"
)

// ProductRepository 商品数据访问层
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository 创建商品仓储实例
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Create 创建新商品
func (r *ProductRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

// GetByID 根据ID获取商品
func (r *ProductRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Update 更新商品信息
func (r *ProductRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

// Delete 删除商品（软删除）
func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}

// List 获取商品列表
func (r *ProductRepository) List(page, pageSize int) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	// 获取总数
	if err := r.db.Model(&model.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page - 1) * pageSize
	err := r.db.Offset(offset).Limit(pageSize).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// DeductStock 扣减库存
func (r *ProductRepository) DeductStock(tx *gorm.DB, productID uint, quantity int) error {
	result := tx.Model(&model.Product{}).
		Where("id = ? AND stock >= ?", productID, quantity).
		UpdateColumn("stock", gorm.Expr("stock - ?", quantity))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrInsufficientStock
	}

	return nil
}
