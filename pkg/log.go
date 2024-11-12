package pkg

import (
	"gin_template/global"

	"go.uber.org/zap"
)

// Info 打印 Info 级别日志
func Info(message string, fields ...zap.Field) {
	global.Logger.Info(message, fields...)
}

// Warn 打印 Warn 级别日志
func Warn(message string, fields ...zap.Field) {
	global.Logger.Warn(message, fields...)
}

// Error 打印 Error 级别日志
func Error(message string, err error, fields ...zap.Field) {
	global.Logger.Error(message, append(fields, zap.Error(err))...)
}

// Debug 打印 Debug 级别日志
func Debug(message string, fields ...zap.Field) {
	global.Logger.Debug(message, fields...)
}

// Fatal 打印 Fatal 级别日志（程序退出）
func Fatal(message string, err error, fields ...zap.Field) {
	global.Logger.Fatal(message, append(fields, zap.Error(err))...)
}
