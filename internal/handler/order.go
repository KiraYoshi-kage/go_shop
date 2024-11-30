package handler

import (
	"myshop/internal/model"
	"myshop/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

// @Summary 创建订单
// @Description 创建新订单
// @Tags 订单管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param order body model.Order true "订单信息"
// @Success 200 {object} map[string]interface{} "创建成功"
// @Failure 400 {object} map[string]interface{} "参数错误"
// @Failure 401 {object} map[string]interface{} "未授权"
// @Failure 500 {object} map[string]interface{} "库存不足"
// @Router /orders [post]
func (h *OrderHandler) Create(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": "参数错误"})
		return
	}

	// 从JWT中获取用户ID
	userID, _ := c.Get("userID")
	order.UserID = userID.(uint)

	if err := h.orderService.Create(c.Request.Context(), &order); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "订单创建成功",
		"data":    order,
	})
}

// @Summary 获取订单详情
// @Description 获取订单详细信息
// @Tags 订单管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "订单ID"
// @Success 200 {object} model.Order
// @Failure 404 {object} map[string]interface{} "订单不存在"
// @Router /orders/{id} [get]
func (h *OrderHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "无效的订单ID"})
		return
	}

	order, err := h.orderService.GetByID(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": "获取订单失败"})
		return
	}

	c.JSON(200, gin.H{"data": order})
}

// @Summary 获取用户订单列表
// @Description 获取当前用户的订单列表
// @Tags 订单管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} ListResponse{data=[]model.Order} "订单列表"
// @Router /orders [get]
func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userID, _ := c.Get("userID")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	orders, total, err := h.orderService.GetUserOrders(userID.(uint), page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取订单列表失败"})
		return
	}

	c.JSON(200, gin.H{
		"data":      orders,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// 添加获取service的方法
func (h *OrderHandler) GetService() *service.OrderService {
	return h.orderService
}
