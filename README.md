# MyShop 电商系统

## 项目简介

MyShop是一个基于Go语言开发的现代电商系统，采用Gin框架构建。本项目实现了基础的电商功能，并集成了Swagger文档系统，便于接口管理和测试。

## 技术栈

- 语言：Go
- Web框架：Gin
- API文档：Swagger
- 数据库：MySQL
- 缓存：Redis

## 项目结构

```
myshop/
├── cmd/
│   └── main.go        # 程序入口
├── docs/
│   └── docs.go        # Swagger文档
├── internal/
│   ├── handler/       # 请求处理器
│   ├── model/         # 数据模型
│   ├── repository/    # 数据访问层
│   └── service/       # 业务逻辑层
└── go.mod             # 依赖管理
```

## 快速开始

### 环境要求

- Go 1.16+
- MySQL 5.7+
- Redis 6.0+

### 安装步骤

1. 克隆项目
```

   git clone https://github.com/你的用户名/myshop.git
   cd myshop
```
2. 安装依赖
```

   go mod tidy
```
3. 配置数据库
```

# 创建配置文件

config.yaml

# 修改配置文件中的数据库连接信息

```

4. 运行项目
   ```

   go run cmd/main.go
   ```

### 访问Swagger文档

启动服务后，访问：http://localhost:8080/swagger/index.html

## 主要功能

- 用户管理
- 商品管理
- 订单处理
- 购物车功能

## 接口文档

详细的API文档请参考Swagger页面，包含：

- 接口描述
- 请求参数
- 响应示例
- 在线测试功能

## 开发指南

1. 代码规范遵循Go标准
2. 提交代码前请运行测试
3. 新功能请先创建分支开发

## 版本历史

- v0.1.0 - 初始版本
  - 基础架构搭建
  - Swagger集成
  - 基本API实现
