package helpers

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.SugaredLogger

func init() {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&lumberjack.Logger{
			Filename:   "chezz.log",
			MaxSize:    10, // megabytes
			MaxBackups: 5,
			MaxAge:     28, // days
		})),
		zap.NewAtomicLevelAt(zap.DebugLevel),
	)
	Logger = zap.New(core, zap.AddCaller()).Sugar()
}
