package handler

import (
	"myshop/internal/model"
	"myshop/internal/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// CreateProductRequest 创建商品请求
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required" example:"iPhone 15"`
	Description string  `json:"description" example:"最新款iPhone"`
	Price       float64 `json:"price" binding:"required,gt=0" example:"6999.00"`
	Stock       int     `json:"stock" binding:"required,gte=0" example:"100"`
}

// ProductResponse 商品响应
type ProductResponse struct {
	ID          uint      `json:"id" example:"1"`
	Name        string    `json:"name" example:"iPhone 15"`
	Description string    `json:"description" example:"最新款iPhone"`
	Price       float64   `json:"price" example:"6999.00"`
	Stock       int       `json:"stock" example:"100"`
	CreatedAt   time.Time `json:"created_at" example:"2023-12-20T10:00:00Z"`
}

// @Summary 创建商品
// @Description 创建新商品（需要管理员权限）
// @Tags 商品管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body CreateProductRequest true "商品信息"
// @Success 200 {object} Response{data=ProductResponse} "创建成功"
// @Failure 400 {object} ErrorResponse "参数错误"
// @Failure 401 {object} ErrorResponse "未授权"
// @Router /products [post]
func (h *ProductHandler) Create(c *gin.Context) {
	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, ErrorResponse{Code: 400, Message: "参数错误"})
		return
	}

	product := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if err := h.productService.Create(product); err != nil {
		c.JSON(500, ErrorResponse{Code: 500, Message: "创建商品失败"})
		return
	}

	c.JSON(200, Response{
		Code:    200,
		Message: "创建成功",
		Data:    product,
	})
}

// @Summary 获取商品列表
// @Description 获取商品列表，支持分页
// @Tags 商品管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} ListResponse{data=[]model.Product} "商品列表"
// @Router /products [get]
func (h *ProductHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	products, total, err := h.productService.List(page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取商品列表失败"})
		return
	}

	c.JSON(200, gin.H{
		"data":      products,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// @Summary 获取商品详情
// @Description 根据ID获取商品详情
// @Tags 商品管理
// @Accept json
// @Produce json
// @Param id path int true "商品ID"
// @Success 200 {object} model.Product
// @Failure 404 {object} map[string]interface{} "商品不存在"
// @Router /products/{id} [get]
func (h *ProductHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "无效的商品ID"})
		return
	}

	product, err := h.productService.GetByID(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": "获取商品失败"})
		return
	}

	c.JSON(200, gin.H{"data": product})
}

// @Summary 更新商品
// @Description 更新商品信息
// @Tags 商品管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "商品ID"
// @Param product body model.Product true "商品信息"
// @Success 200 {object} map[string]interface{} "更新成功"
// @Failure 400,404 {object} map[string]interface{} "参数错误或商品不存在"
// @Router /products/{id} [put]
func (h *ProductHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "无效的商品ID"})
		return
	}

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": "参数错误"})
		return
	}

	product.ID = uint(id)
	if err := h.productService.Update(&product); err != nil {
		c.JSON(500, gin.H{"error": "更新商品失败"})
		return
	}

	c.JSON(200, gin.H{"message": "更新成功"})
}

// @Summary 删除商品
// @Description 删除指定商品
// @Tags 商品管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "商品ID"
// @Success 200 {object} map[string]interface{} "删除成功"
// @Failure 404 {object} map[string]interface{} "商品不存在"
// @Router /products/{id} [delete]
func (h *ProductHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "无效的商品ID"})
		return
	}

	if err := h.productService.Delete(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": "删除商品失败"})
		return
	}

	c.JSON(200, gin.H{"message": "删除成功"})
}

// ListResponse represents a generic list response
type ListResponse struct {
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}
