package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
// 采用GORM标签定义数据库表结构
type User struct {
	ID        uint           `gorm:"primarykey"`          // 用户ID，主键
	Username  string         `gorm:"uniqueIndex;size:32"` // 用户名，唯一索引，最大长度32
	Password  string         `gorm:"size:128" json:"-"`   // 密码，最大长度128，json序列化时忽略
	CreatedAt time.Time      // 创建时间，GORM自动维护
	UpdatedAt time.Time      // 更新时间，GORM自动维护
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 软删除时间，支持软删除
}
