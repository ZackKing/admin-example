package log

import (
	"fmt"
	"os"
	"runtime"

	"admin-api/internal/services/conf"
	"admin-api/pkg/lib/def"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger(conf *conf.Log) *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.CallerKey = ""
	encoder := zapcore.NewJSONEncoder(config)

	level := zap.NewAtomicLevel()
	switch conf.Level {
	case "info":
		level.SetLevel(zap.InfoLevel)
	case "error":
		level.SetLevel(zap.ErrorLevel)
	default:
		level.SetLevel(zap.DebugLevel)
	}

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(logfile), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return true
		})),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapcore.ErrorLevel
		})),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stderr), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})),
	)

	logger := zap.New(core, zap.AddCaller())
	return logger
}

func changeZapFields(m def.KV) []zapcore.Field {
	fields := []zapcore.Field{}
	_, file, line, _ := runtime.Caller(2)
	fields = append(fields, zap.String("file", fmt.Sprintf("%s:%d", file, line)))
	for k, v := range m {
		fields = append(fields, zap.Any(k, v))
	}
	return fields
}
