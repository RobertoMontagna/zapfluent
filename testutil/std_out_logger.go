package testutil

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/pkg/core"
)

// StdOutLogger creates a new zap.SugaredLogger that is configured to write
// JSON-formatted logs to standard output.
func StdOutLogger() *zap.SugaredLogger {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	coreEncoder := zapcore.NewCore(
		core.NewFluentEncoder(
			zapcore.NewJSONEncoder(encoderCfg),
			core.NewConfiguration(),
		),
		os.Stdout,
		zap.DebugLevel,
	)
	logger := zap.New(coreEncoder)
	zap.ReplaceGlobals(logger)
	return zap.S()
}
