package service

import (
	"context"
	"fmt"
	"myshop/internal/model"
	"myshop/internal/repository"
	"time"

	"gorm.io/gorm"
)

type OrderService struct {
	orderRepo   *repository.OrderRepository
	productRepo *repository.ProductRepository
}

func NewOrderService(orderRepo *repository.OrderRepository, productRepo *repository.ProductRepository) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *OrderService) Create(ctx context.Context, order *model.Order) error {
	order.OrderNo = fmt.Sprintf("%d%d", time.Now().UnixNano(), order.UserID)
	order.Status = model.OrderStatusPending

	var totalPrice float64
	for _, item := range order.Items {
		product, err := s.productRepo.GetByID(item.ProductID)
		if err != nil {
			return fmt.Errorf("获取商品信息失败: %w", err)
		}
		item.Price = product.Price
		totalPrice += product.Price * float64(item.Quantity)
	}
	order.TotalPrice = totalPrice

	err := s.orderRepo.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return fmt.Errorf("创建订单失败: %w", err)
		}

		for _, item := range order.Items {
			if err := s.productRepo.DeductStock(tx, item.ProductID, item.Quantity); err != nil {
				return fmt.Errorf("扣减库存失败: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *OrderService) GetByID(id uint) (*model.Order, error) {
	return s.orderRepo.GetByID(id)
}

func (s *OrderService) GetUserOrders(userID uint, page, pageSize int) ([]model.Order, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	return s.orderRepo.GetByUserID(userID, page, pageSize)
}

func (s *OrderService) UpdateStatus(id uint, status int) error {
	return s.orderRepo.UpdateStatus(id, status)
}
