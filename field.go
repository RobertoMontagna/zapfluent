package zapfluent

import "go.uber.org/zap/zapcore"

type Field interface {
	Name() string
	Encode(zapcore.ObjectEncoder) error
}
