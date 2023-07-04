package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Slogger *zap.SugaredLogger

// log to console
func init() {
	// 创建一个日志配置
	config := zap.NewDevelopmentConfig()

	// 设置日志级别为 DEBUG
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)

	// 将编码器配置为控制台编码器
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 在控制台使用带颜色的级别标签
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // 使用 ISO8601 格式化时间

	// 创建 SugarLogger
	logger, _ := config.Build(zap.AddCaller(), zap.AddCallerSkip(1), zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(core, zapcore.NewCore(
			zapcore.NewConsoleEncoder(config.EncoderConfig),
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel,
		))
	}))

	Slogger = logger.Sugar()
}

// format: "debug: xxxx%s", url
func DebugF(template string, args ...interface{}) {
	Slogger.Debugf(template, args...)
	//Slogger.Infow()
}

func ErrorF(template string, args ...interface{}) {
	Slogger.Errorf(template, args...)
}

func InfoF(template string, args ...interface{}) {
	Slogger.Infof(template, args...)
}
