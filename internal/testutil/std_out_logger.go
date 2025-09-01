package testutil

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/pkg/core"
)

// StdoutLoggerForTest creates and returns a sugared zap logger configured for tests that
// writes JSON to stdout.
// The encoder uses "msg" for messages, "level" for levels and "logger" for logger name, encodes
// levels lowercase, times in ISO8601 and durations as strings. It replaces the global logger via
// zap.ReplaceGlobals before returning the sugared logger.
func StdoutLoggerForTest(options ...zap.Option) *zap.SugaredLogger {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		TimeKey:        "time",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	coreEncoder := zapcore.NewCore(
		core.NewFluentEncoder(
			zapcore.NewJSONEncoder(encoderCfg),
			core.NewConfiguration(),
		),
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel,
	)
	logger := zap.New(coreEncoder, options...)
	zap.ReplaceGlobals(logger)
	return zap.S()
}
