package handler

// Response 通用响应结构
type Response struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

// PageResponse 分页响应结构
type PageResponse struct {
	Code     int         `json:"code" example:"200"`
	Message  string      `json:"message" example:"success"`
	Data     interface{} `json:"data"`
	Total    int64       `json:"total" example:"100"`
	Page     int         `json:"page" example:"1"`
	PageSize int         `json:"page_size" example:"10"`
}

// ErrorResponse 错误响应结构
type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"参数错误"`
}
