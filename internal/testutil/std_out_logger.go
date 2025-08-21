package testutil

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
)

// StdOutLogger creates a new zap.SugaredLogger that is configured to write
// JSON-formatted logs to standard output.
//
// This is useful for debugging and running examples.
func StdOutLogger() *zap.SugaredLogger {
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
			zapfluent.NewConfiguration(),
		),
		os.Stdout,
		zap.DebugLevel,
	)
	logger := zap.New(core)

	zap.ReplaceGlobals(logger)

	return zap.S()
}
