package zaptestutil

import "time"

// ConstantClockForTest is a deterministic zapcore.Clock implementation for tests.
// Now always returns the fixed time. NewTicker returns a zero-value ticker
// (non-ticking, nil channel) to avoid goroutines and ensure determinism.
type ConstantClockForTest time.Time

// NewConstantClockForTest constructs a ConstantClockForTest from t.
func NewConstantClockForTest(t time.Time) ConstantClockForTest {
	return ConstantClockForTest(t)
}

func (c ConstantClockForTest) Now() time.Time {
	return time.Time(c)
}

func (c ConstantClockForTest) NewTicker(_ time.Duration) *time.Ticker {
	return &time.Ticker{
		C: nil,
	}
}
