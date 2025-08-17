package zapfluent_test

import (
	"os"

	"go.robertomontagna.dev/zapfluent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
			zapfluent.DefaultFluentConfig(),
		),
		os.Stdout,
		zap.DebugLevel,
	)
	logger := zap.New(core)

	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	return zap.S()
}
