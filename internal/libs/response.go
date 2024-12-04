package libs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standardized response structure for all API responses.
type Response struct {
	Code    int         `json:"code"`            // 业务状态码：0 为成功，-1 为失败
	Message string      `json:"message"`         // 消息内容
	Success bool        `json:"success"`         // 是否成功
	Data    interface{} `json:"data"`            // 返回数据
	Extra   interface{} `json:"extra,omitempty"` // 可选的扩展数据
}

// 自定义的业务状态码
const (
	CodeSuccess = 0
	CodeFail    = -1
)

// NewResponse 构造通用响应
func NewResponse(code int, message string, success bool, data interface{}, extra interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Success: success,
		Data:    data,
		Extra:   extra,
	}
}

// sendResponse 通过 ctx 返回 JSON 响应，同时设置 HTTP 状态码
func sendResponse(ctx *gin.Context, httpCode int, response Response) {
	ctx.AbortWithStatusJSON(httpCode, response)
}

// SuccessResponse 返回成功的响应
func SuccessResponse(ctx *gin.Context, message string, data interface{}) {
	response := NewResponse(CodeSuccess, message, true, data, nil)
	sendResponse(ctx, http.StatusOK, response)
}

// CreatedResponse 返回资源创建成功的响应
func CreatedResponse(ctx *gin.Context, message string, data interface{}) {
	response := NewResponse(CodeSuccess, message, true, data, nil)
	sendResponse(ctx, http.StatusCreated, response)
}

// FailResponse 返回失败的响应
func FailResponse(ctx *gin.Context, message string, data interface{}) {
	response := NewResponse(CodeFail, message, false, data, nil)
	sendResponse(ctx, http.StatusBadRequest, response)
}

// NotFoundResponse 返回未找到的响应
func NotFoundResponse(ctx *gin.Context, message string) {
	response := NewResponse(CodeFail, message, false, nil, nil)
	sendResponse(ctx, http.StatusNotFound, response)
}

// UnauthorizedResponse 返回未授权的响应
func UnauthorizedResponse(ctx *gin.Context, message string) {
	response := NewResponse(CodeFail, message, false, nil, nil)
	sendResponse(ctx, http.StatusUnauthorized, response)
}

// ForbiddenResponse 返回禁止访问的响应
func ForbiddenResponse(ctx *gin.Context, message string) {
	response := NewResponse(CodeFail, message, false, nil, nil)
	sendResponse(ctx, http.StatusForbidden, response)
}

// BadRequestResponse 返回请求无效的响应
func BadRequestResponse(ctx *gin.Context, message string) {
	response := NewResponse(CodeFail, message, false, nil, nil)
	sendResponse(ctx, http.StatusBadRequest, response)
}

// InternalServerErrorResponse 返回服务器错误的响应
func InternalServerErrorResponse(ctx *gin.Context, message string) {
	response := NewResponse(CodeFail, message, false, nil, nil)
	sendResponse(ctx, http.StatusInternalServerError, response)
}

// ServiceUnavailableResponse 返回服务不可用的响应
func ServiceUnavailableResponse(ctx *gin.Context, message string) {
	response := NewResponse(CodeFail, message, false, nil, nil)
	sendResponse(ctx, http.StatusServiceUnavailable, response)
}

// AddExtra 可以向现有响应添加额外信息
func AddExtra(ctx *gin.Context, response Response, extra interface{}) {
	response.Extra = extra
	sendResponse(ctx, http.StatusOK, response)
}
