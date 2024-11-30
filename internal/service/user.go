package service

import (
	"myshop/internal/model"
	"myshop/internal/repository"
	"myshop/pkg/utils"
)

// UserService 用户业务逻辑层
type UserService struct {
	repo *repository.UserRepository // 用户数据仓储
}

// NewUserService 创建用户服务实例
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Register 用户注册
// 1. 检查用户名是否已存在
// 2. 对密码进行加密
// 3. 创建新用户
func (s *UserService) Register(user *model.User) error {
	// 检查用户名是否已存在
	existingUser, err := s.repo.GetByUsername(user.Username)
	if err == nil && existingUser != nil {
		return ErrUserExists
	}

	// 密码加密
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// 创建用户
	return s.repo.Create(user)
}

// Login 用户登录
// 1. 根据用户名查找用户
// 2. 验证密码
// 3. 生成JWT token
func (s *UserService) Login(username, password string) (string, error) {
	// 查找用户
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	// 验证密码
	if !utils.CheckPassword(password, user.Password) {
		return "", ErrInvalidCredentials
	}

	// 生成token
	return utils.GenerateToken(user.ID)
}

// GetByID 根据ID获取用户信息
func (s *UserService) GetByID(id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}
