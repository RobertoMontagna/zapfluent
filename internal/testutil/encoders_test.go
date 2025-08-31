package testutil_test

import (
	"testing"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/internal/testutil"

	. "github.com/onsi/gomega"
)

func TestNewDoNotEncodeEncoder(t *testing.T) {
	g := NewWithT(t)

	enc := zapcore.NewMapObjectEncoder()

	result := testutil.NewDoNotEncodeEncoderForTest(enc)

	g.Expect(result.ObjectEncoder).To(Equal(enc))
}

func TestDoNotEncodeEncoder_Clone(t *testing.T) {
	g := NewWithT(t)

	sut := testutil.NewDoNotEncodeEncoderForTest(zapcore.NewMapObjectEncoder())

	clone := sut.Clone()

	g.Expect(clone).ToNot(BeIdenticalTo(sut))
}

func TestDoNotEncodeEncoder_EncodeEntry(t *testing.T) {
	g := NewWithT(t)

	sut := testutil.NewDoNotEncodeEncoderForTest(zapcore.NewMapObjectEncoder())

	buffer, err := sut.EncodeEntry(zapcore.Entry{}, nil)

	g.Expect(buffer).To(BeNil())
	g.Expect(err).ToNot(HaveOccurred())
}
