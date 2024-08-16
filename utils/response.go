package utils

type Response struct {
	Code    MyCode      `json:"code"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
type MyCode int

const (
	// SUCCESS 正常
	SUCCESS MyCode = 0
	// ERROR 参数，前端传入等异常
	ERROR MyCode = 1000
	// BAD 服务器内部出现问题
	BAD MyCode = 5000
)

func OK(message string, data interface{}) Response {
	return Response{
		Code:    SUCCESS,
		Message: message,
		Success: true,
		Data:    data,
	}
}

func OKWithMessage(message string) Response {
	return Response{
		Code:    SUCCESS,
		Message: message,
		Success: true,
		Data:    nil,
	}
}

func OKWithData(data interface{}) Response {
	return Response{
		Code:    SUCCESS,
		Message: "success",
		Success: true,
		Data:    data,
	}
}

func Fail(message string, data interface{}) Response {
	return Response{
		Code:    ERROR,
		Message: message,
		Success: false,
		Data:    data,
	}
}
func FailWithMessage(message string) Response {
	return Response{
		Code:    ERROR,
		Message: message,
		Success: false,
		Data:    nil,
	}
}
func FailWithDate(data interface{}) Response {
	return Response{
		Code:    ERROR,
		Message: "请求出错，请检查",
		Success: false,
		Data:    data,
	}
}
func Bad(message string, data interface{}) Response {
	return Response{
		Code:    BAD,
		Message: message,
		Success: false,
		Data:    data,
	}
}
func BadWithMessage(message string) Response {
	return Response{
		Code:    BAD,
		Message: message,
		Success: false,
		Data:    nil,
	}
}
func BadWithDate(data interface{}) Response {
	return Response{
		Code:    BAD,
		Message: "服务器内部出现错误，请联系开发者",
		Success: false,
		Data:    data,
	}
}
