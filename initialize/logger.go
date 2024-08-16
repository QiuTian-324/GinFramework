package initialize

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// LoggerInitialize 日志初始化
func LoggerInitialize() *zap.SugaredLogger {
	logMode := zapcore.InfoLevel
	if viper.GetBool("development.develop") {
		logMode = zapcore.DebugLevel
	}
	core := zapcore.NewCore(
		getEncoder(),
		zapcore.NewMultiWriteSyncer(getLogWriter()),
		// 同时在控制台和文件显示日志
		//zapcore.NewMultiWriteSyncer(getLogWriter(), zapcore.AddSync(os.Stdout))
		logMode)

	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
func getLogWriter() zapcore.WriteSyncer {
	stSeparator := string(os.PathSeparator)
	stRootDir, _ := os.Getwd()
	filePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly) + ".log"
	// 分割器
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    viper.GetInt("log.MaxSize"), // 最大M数，超过则分割
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"), // 最大保留天数
		Compress:   false,                      // 是否压缩 disabled by default
	}
	return zapcore.AddSync(lumberJackLogger)
}
