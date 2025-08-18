package fluentfield_test

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
)

func stdOutLogger() *zap.SugaredLogger {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	core := zapcore.NewCore(
		zapfluent.NewFluentEncoder(
			zapcore.NewJSONEncoder(encoderCfg),
			zapfluent.NewFluentConfig(),
		),
		os.Stdout,
		zap.DebugLevel,
	)
	logger := zap.New(core)

	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	return zap.S()
}
