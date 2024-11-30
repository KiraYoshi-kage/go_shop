package handler

import (
	"myshop/internal/model"
	"myshop/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// RegisterRequest 注册请求结构
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=32" example:"testuser"`
	Password string `json:"password" binding:"required,min=6,max=32" example:"password123"`
}

// RegisterResponse 注册响应结构
type RegisterResponse struct {
	Message string `json:"message" example:"注册成功"`
	UserID  uint   `json:"user_id" example:"1"`
}

// @Summary 用户注册
// @Description 创建新用户账号（仅需用户名和密码）
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "用户名和密码"
// @Success 200 {object} RegisterResponse
// @Failure 400 {object} ErrorResponse "参数错误或用户名已存在"
// @Failure 500 {object} ErrorResponse "服务器错误"
// @Router /user/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, ErrorResponse{Message: "参数错误: 用户名长度3-32位，密码长度6-32位"})
		return
	}

	user := &model.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := h.userService.Register(user); err != nil {
		if err == service.ErrUserExists {
			c.JSON(400, ErrorResponse{Message: "用户名已存在"})
			return
		}
		c.JSON(500, ErrorResponse{Message: "注册失败"})
		return
	}

	c.JSON(200, RegisterResponse{
		Message: "注册成功",
		UserID:  user.ID,
	})
}

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"testuser"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// LoginResponse 登录响应结构
type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIs..."`
}

// @Summary 用户登录
// @Description 使用用户名和密码登录获取token
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body LoginRequest true "用户名和密码"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse "参数错误"
// @Failure 401 {object} ErrorResponse "用户名或密码错误"
// @Router /user/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, ErrorResponse{Message: "参数错误"})
		return
	}

	token, err := h.userService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(401, ErrorResponse{Message: "用户名或密码错误"})
		return
	}

	c.JSON(200, LoginResponse{Token: token})
}

// UserInfo 用户信息响应结构
type UserInfo struct {
	ID       uint   `json:"id" example:"1"`
	Username string `json:"username" example:"testuser"`
}

// @Summary 获取用户信息
// @Description 获取当前登录用户的信息（仅返回ID和用户名）
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} UserInfo
// @Failure 401 {object} ErrorResponse "未授权"
// @Failure 500 {object} ErrorResponse "服务器错误"
// @Router /user/info [get]
func (h *UserHandler) GetInfo(c *gin.Context) {
	userID, _ := c.Get("userID")
	user, err := h.userService.GetByID(userID.(uint))
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "获取用户信息失败"})
		return
	}

	c.JSON(200, UserInfo{
		ID:       user.ID,
		Username: user.Username,
	})
}
