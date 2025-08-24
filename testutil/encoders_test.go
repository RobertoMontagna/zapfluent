package testutil_test

import (
	"testing"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/testutil"

	. "github.com/onsi/gomega"
)

func TestDoNotEncodeEncoder_WrapsObjectEncoder(t *testing.T) {
	g := NewWithT(t)

	mapEncoder := zapcore.NewMapObjectEncoder()
	testEncoder := testutil.NewDoNotEncodeEncoder(mapEncoder)

	testEncoder.AddString("key", "value")

	g.Expect(mapEncoder.Fields).To(HaveKeyWithValue("key", "value"))
}

func TestDoNotEncodeEncoder_CloneReturnsNewInstance(t *testing.T) {
	g := NewWithT(t)

	mapEncoder := zapcore.NewMapObjectEncoder()
	testEncoder := testutil.NewDoNotEncodeEncoder(mapEncoder)

	cloned := testEncoder.Clone()

	g.Expect(cloned).ToNot(BeIdenticalTo(testEncoder))
	g.Expect(cloned).To(BeAssignableToTypeOf(&testutil.DoNotEncodeEncoder{}))
}

func TestDoNotEncodeEncoder_EncodeEntryIsNoOp(t *testing.T) {
	g := NewWithT(t)

	testEncoder := testutil.NewDoNotEncodeEncoder(zapcore.NewMapObjectEncoder())

	buf, err := testEncoder.EncodeEntry(zapcore.Entry{}, nil)

	g.Expect(buf).To(BeNil())
	g.Expect(err).ToNot(HaveOccurred())
}
