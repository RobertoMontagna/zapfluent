package zapfluent

import "go.uber.org/zap/zapcore"

type Field interface {
	Encode(zapcore.ObjectEncoder) error
}
