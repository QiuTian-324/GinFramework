package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// LoggerInternal 初始化日志系统
func LoggerInternal() *zap.Logger {
	logLevel := zapcore.InfoLevel
	if viper.GetBool("development.develop") {
		logLevel = zapcore.DebugLevel
	}

	core := zapcore.NewCore(
		getEncoder(),
		zapcore.NewMultiWriteSyncer(getLogWriter(), zapcore.AddSync(os.Stdout)),
		logLevel,
	)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 不显示调用栈信息
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 彩色日志等级
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder      // 显示简短调用路径

	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter 日志文件配置
func getLogWriter() zapcore.WriteSyncer {
	stSeparator := string(os.PathSeparator)
	stRootDir, _ := os.Getwd()
	filePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format("2006-01-02") + ".log"

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    viper.GetInt("log.max_size"),    // 日志文件最大 MB 数
		MaxBackups: viper.GetInt("log.max_backups"), // 最多保留的旧日志文件数
		MaxAge:     viper.GetInt("log.max_age"),     // 日志文件最多保存的天数
		Compress:   viper.GetBool("log.compress"),   // 是否压缩日志
	}
	return zapcore.AddSync(lumberJackLogger)
}
