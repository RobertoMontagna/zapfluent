package zaptestutil_test

import (
	"testing"
	"time"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/internal/testutil/zaptestutil"

	. "github.com/onsi/gomega"
)

func TestConstantClockForTest_Now(t *testing.T) {
	g := NewWithT(t)

	expectedTime := time.Now().Truncate(time.Second)

	clock := zaptestutil.NewConstantClockForTest(expectedTime)
	actualTime := clock.Now()

	g.Expect(actualTime).To(Equal(expectedTime))
}

func TestConstantClockForTest_NewTicker(t *testing.T) {
	g := NewWithT(t)

	clock := zaptestutil.NewConstantClockForTest(time.Now())

	ticker := clock.NewTicker(1 * time.Minute)

	g.Expect(ticker).ToNot(BeNil())
	g.Expect(ticker.C).To(BeNil())
}

func TestConstantClockForTestCompatibility(t *testing.T) {
	// Compile-time assertion: ConstantClockForTest implements zapcore.Clock.
	var _ zapcore.Clock = zaptestutil.ConstantClockForTest{}
}
