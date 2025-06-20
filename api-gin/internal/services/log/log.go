package log

import (
	"log"
	"sync"

	"admin-api/internal/services/conf"
	"admin-api/pkg/lib/def"

	"go.uber.org/zap"
)

var (
	Logger  *zap.Logger
	logfile *RotateFile
	once    sync.Once
)

func InitLogger(conf conf.Log) {
	once.Do(func() {
		logfile = NewRotateFile(conf.Path)
		Logger = newZapLogger(&conf)
		if Logger == nil {
			panic("Logger init error")
		}
		log.Default().Printf("Logger initialized")
	})
}

func Debug(msg string, data def.KV) {
	fields := changeZapFields(data)
	Logger.Debug(msg, fields...)
}

func Info(msg string, data def.KV) {
	fields := changeZapFields(data)
	Logger.Info(msg, fields...)
}

func Error(msg string, err error, data def.KV) {
	fields := changeZapFields(data)
	fields = append(fields, zap.Error(err))
	Logger.Error(msg, fields...)
}
