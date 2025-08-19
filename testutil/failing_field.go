package testutil

import "go.uber.org/zap/zapcore"

type FailingField struct {
	Err       error
	NameValue string
}

func (f FailingField) Encode(enc zapcore.ObjectEncoder) error {
	return f.Err
}

func (f FailingField) Name() string {
	if f.NameValue == "" {
		return "error"
	}
	return f.NameValue
}
