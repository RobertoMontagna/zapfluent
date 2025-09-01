package zaptestutil_test

import (
	"testing"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/internal/testutil/zaptestutil"
)

func TestConstantClockForTestCompatibility(t *testing.T) {
	// Compile-time assertion: ConstantClockForTest implements zapcore.Clock.
	var _ zapcore.Clock = zaptestutil.ConstantClockForTest{}
}
