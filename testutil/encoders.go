// Package testutil provides helpers for testing purposes.
package testutil

import (
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

// DoNotEncodeEncoder is a wrapper for a zapcore.ObjectEncoder that satisfies the
// zapcore.Encoder interface for testing purposes. It provides no-op
// implementations for the methods that are part of the Encoder interface but
// not the ObjectEncoder interface (Clone and EncodeEntry).
//
// This is useful when you need to pass an encoder to a function that requires a
// full zapcore.Encoder, but you want to use a zapcore.MapObjectEncoder to
// inspect the fields that were logged.
type DoNotEncodeEncoder struct {
	zapcore.ObjectEncoder
}

// NewDoNotEncodeEncoder creates a new DoNotEncodeEncoder that wraps the given
// zapcore.ObjectEncoder.
func NewDoNotEncodeEncoder(enc zapcore.ObjectEncoder) *DoNotEncodeEncoder {
	return &DoNotEncodeEncoder{ObjectEncoder: enc}
}

// Clone provides a fake implementation of Clone to satisfy the Encoder interface.
// It does not produce a true clone.
func (t *DoNotEncodeEncoder) Clone() zapcore.Encoder {
	// The underlying ObjectEncoder cannot be reliably cloned, so we return a
	// new instance. This is acceptable for the test cases that use this utility.
	return &DoNotEncodeEncoder{ObjectEncoder: zapcore.NewMapObjectEncoder()}
}

// EncodeEntry is a no-op implementation to satisfy the Encoder interface.
// The tests that use this wrapper do not rely on the full encoding pipeline,
// only on the fields added to the underlying ObjectEncoder.
func (t *DoNotEncodeEncoder) EncodeEntry(_ zapcore.Entry, _ []zapcore.Field) (*buffer.Buffer, error) {
	return nil, nil
}
