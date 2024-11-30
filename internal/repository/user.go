package repository

import (
	"myshop/internal/model"

	"gorm.io/gorm"
)

// UserRepository 用户数据访问层
// 负责用户相关的数据库操作
type UserRepository struct {
	db *gorm.DB // 数据库连接
}

// NewUserRepository 创建用户仓储实例
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create 创建新用户
func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// GetByUsername 根据用户名查询用户
func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByID 根据ID查询用户
func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
