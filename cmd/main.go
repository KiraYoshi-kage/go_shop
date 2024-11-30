package main

import (
	"fmt"
	"log"
	_ "myshop/docs" // 导入swagger文档
	"myshop/internal/config"
	"myshop/internal/handler"
	"myshop/internal/model"
	"myshop/internal/repository"
	"myshop/internal/service"
	"myshop/pkg/middleware"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files" // 修改这行
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title MyShop API
// @version 1.0
// @description MyShop电商系统API文档
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description 在请求头中添加 Authorization: Bearer {token} 进行身份验证
func main() {
	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/myshop?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 自动迁移数据库表
	err = db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Order{},
		&model.OrderItem{},
	)
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	// 初始化各层依赖
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, productRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	// 初始化路由
	r := gin.Default()

	// API路由
	api := r.Group("/api")
	{
		// 用户相关路由
		api.POST("/user/register", userHandler.Register)
		api.POST("/user/login", userHandler.Login)

		// 商品相关路由
		api.GET("/products", productHandler.List)
		api.GET("/products/:id", productHandler.GetByID)

		// 需要认证的路由
		auth := api.Group("/", middleware.Auth())
		{
			// 用户
			auth.GET("/user/info", userHandler.GetInfo)

			// 商品管理
			auth.POST("/products", productHandler.Create)
			auth.PUT("/products/:id", productHandler.Update)
			auth.DELETE("/products/:id", productHandler.Delete)

			// 订单管理
			auth.POST("/orders", orderHandler.Create)
			auth.GET("/orders/:id", orderHandler.GetByID)
			auth.GET("/orders", orderHandler.GetUserOrders)
		}
	}
	// 添加swagger路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	// 加载配置
	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 启动服务器
	if err := r.Run(fmt.Sprintf(":%d", config.Server.Port)); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
