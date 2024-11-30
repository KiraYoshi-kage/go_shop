package service

import (
	"myshop/internal/model"
	"myshop/internal/repository"
)

// ProductService 商品业务逻辑层
type ProductService struct {
	repo *repository.ProductRepository // 商品仓储
}

// NewProductService 创建商品服务实例
func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// Create 创建新商品
func (s *ProductService) Create(product *model.Product) error {
	return s.repo.Create(product)
}

// GetByID 根据ID获取商品
func (s *ProductService) GetByID(id uint) (*model.Product, error) {
	return s.repo.GetByID(id)
}

// Update 更新商品信息
func (s *ProductService) Update(product *model.Product) error {
	return s.repo.Update(product)
}

// Delete 删除商品
func (s *ProductService) Delete(id uint) error {
	return s.repo.Delete(id)
}

// List 获取商品列表
func (s *ProductService) List(page, pageSize int) ([]model.Product, int64, error) {
	// 参数验证
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	return s.repo.List(page, pageSize)
}
