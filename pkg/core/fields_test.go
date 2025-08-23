package core_test

import (
	"testing"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/pkg/core"

	. "github.com/onsi/gomega"
)

// MARK: Test Structs
// =============================================================================

type testObject struct {
	value string
}

func (t *testObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("value", t.value)
	return nil
}

type testComparableObject struct {
	value string
}

func (t testComparableObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("value", t.value)
	return nil
}

// MARK: Tests
// =============================================================================

func TestString(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)

	t.Run("it creates a string field correctly", func(t *testing.T) {
		// Arrange
		field := core.String("my-key", "my-value")
		enc := zapcore.NewMapObjectEncoder()

		// Act
		err := field.Encode(enc)

		// Assert
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(enc.Fields).To(HaveKeyWithValue("my-key", "my-value"))
		g.Expect(field.Name()).To(Equal("my-key"))
	})

	t.Run("NonZero filter works correctly", func(t *testing.T) {
		// Arrange
		zeroField := core.String("zero-key", "").NonZero()
		nonZeroField := core.String("non-zero-key", "value").NonZero()
		enc := zapcore.NewMapObjectEncoder()

		// Act
		err1 := zeroField.Encode(enc)
		err2 := nonZeroField.Encode(enc)

		// Assert
		g.Expect(err1).ToNot(HaveOccurred())
		g.Expect(err2).ToNot(HaveOccurred())
		g.Expect(enc.Fields).ToNot(HaveKey("zero-key"))
		g.Expect(enc.Fields).To(HaveKeyWithValue("non-zero-key", "value"))
	})
}

func TestInt(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)

	t.Run("it creates an int field correctly", func(t *testing.T) {
		// Arrange
		field := core.Int("my-key", 123)
		enc := zapcore.NewMapObjectEncoder()

		// Act
		err := field.Encode(enc)

		// Assert
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(enc.Fields).To(HaveKeyWithValue("my-key", 123))
		g.Expect(field.Name()).To(Equal("my-key"))
	})

	t.Run("NonZero filter works correctly", func(t *testing.T) {
		// Arrange
		zeroField := core.Int("zero-key", 0).NonZero()
		nonZeroField := core.Int("non-zero-key", 42).NonZero()
		enc := zapcore.NewMapObjectEncoder()

		// Act
		err1 := zeroField.Encode(enc)
		err2 := nonZeroField.Encode(enc)

		// Assert
		g.Expect(err1).ToNot(HaveOccurred())
		g.Expect(err2).ToNot(HaveOccurred())
		g.Expect(enc.Fields).ToNot(HaveKey("zero-key"))
		g.Expect(enc.Fields).To(HaveKeyWithValue("non-zero-key", 42))
	})
}

func TestInt8(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)

	t.Run("it creates an int8 field correctly", func(t *testing.T) {
		// Arrange
		field := core.Int8("my-key", 12)
		enc := zapcore.NewMapObjectEncoder()

		// Act
		err := field.Encode(enc)

		// Assert
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(enc.Fields).To(HaveKeyWithValue("my-key", int8(12)))
		g.Expect(field.Name()).To(Equal("my-key"))
	})

	t.Run("NonZero filter works correctly", func(t *testing.T) {
		// Arrange
		zeroField := core.Int8("zero-key", 0).NonZero()
		nonZeroField := core.Int8("non-zero-key", 4).NonZero()
		enc := zapcore.NewMapObjectEncoder()

		// Act
		err1 := zeroField.Encode(enc)
		err2 := nonZeroField.Encode(enc)

		// Assert
		g.Expect(err1).ToNot(HaveOccurred())
		g.Expect(err2).ToNot(HaveOccurred())
		g.Expect(enc.Fields).ToNot(HaveKey("zero-key"))
		g.Expect(enc.Fields).To(HaveKeyWithValue("non-zero-key", int8(4)))
	})
}

func TestObject(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)

	t.Run("it creates an object field correctly", func(t *testing.T) {
		// Arrange
		obj := &testObject{value: "test"}
		field := core.Object("my-key", obj, func(o *testObject) bool { return o != nil && o.value != "" })
		enc := zapcore.NewMapObjectEncoder()

		// Act
		err := field.Encode(enc)

		// Assert
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(enc.Fields).To(HaveKey("my-key"))
		g.Expect(enc.Fields["my-key"]).To(Equal(map[string]interface{}{"value": "test"}))
		g.Expect(field.Name()).To(Equal("my-key"))
	})

	t.Run("NonZero filter works correctly", func(t *testing.T) {
		// Arrange
		isNonZero := func(o *testObject) bool { return o != nil && o.value != "" }
		zeroField := core.Object("zero-key", &testObject{value: ""}, isNonZero).NonZero()
		nonZeroField := core.Object("non-zero-key", &testObject{value: "value"}, isNonZero).NonZero()
		enc := zapcore.NewMapObjectEncoder()

		// Act
		err1 := zeroField.Encode(enc)
		err2 := nonZeroField.Encode(enc)

		// Assert
		g.Expect(err1).ToNot(HaveOccurred())
		g.Expect(err2).ToNot(HaveOccurred())
		g.Expect(enc.Fields).ToNot(HaveKey("zero-key"))
		g.Expect(enc.Fields).To(HaveKey("non-zero-key"))
	})
}

func TestBool(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)

	t.Run("it creates a bool field correctly", func(t *testing.T) {
		// Arrange
		field := core.Bool("my-key", true)
		enc := zapcore.NewMapObjectEncoder()

		// Act
		err := field.Encode(enc)

		// Assert
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(enc.Fields).To(HaveKeyWithValue("my-key", true))
		g.Expect(field.Name()).To(Equal("my-key"))
	})

	t.Run("NonZero filter works correctly", func(t *testing.T) {
		// Arrange
		zeroField := core.Bool("zero-key", false).NonZero()
		nonZeroField := core.Bool("non-zero-key", true).NonZero()
		enc := zapcore.NewMapObjectEncoder()

		// Act
		err1 := zeroField.Encode(enc)
		err2 := nonZeroField.Encode(enc)

		// Assert
		g.Expect(err1).ToNot(HaveOccurred())
		g.Expect(err2).ToNot(HaveOccurred())
		g.Expect(enc.Fields).ToNot(HaveKey("zero-key"))
		g.Expect(enc.Fields).To(HaveKeyWithValue("non-zero-key", true))
	})
}

func TestComparableObject(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)

	t.Run("it creates a comparable object field correctly", func(t *testing.T) {
		// Arrange
		obj := testComparableObject{value: "test"}
		field := core.ComparableObject("my-key", obj)
		enc := zapcore.NewMapObjectEncoder()

		// Act
		err := field.Encode(enc)

		// Assert
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(enc.Fields).To(HaveKey("my-key"))
		g.Expect(enc.Fields["my-key"]).To(Equal(map[string]interface{}{"value": "test"}))
		g.Expect(field.Name()).To(Equal("my-key"))
	})

	t.Run("NonZero filter works correctly", func(t *testing.T) {
		// Arrange
		var zeroValue testComparableObject
		zeroField := core.ComparableObject("zero-key", zeroValue).NonZero()
		nonZeroField := core.ComparableObject("non-zero-key", testComparableObject{value: "value"}).NonZero()
		enc := zapcore.NewMapObjectEncoder()

		// Act
		err1 := zeroField.Encode(enc)
		err2 := nonZeroField.Encode(enc)

		// Assert
		g.Expect(err1).ToNot(HaveOccurred())
		g.Expect(err2).ToNot(HaveOccurred())
		g.Expect(enc.Fields).ToNot(HaveKey("zero-key"))
		g.Expect(enc.Fields).To(HaveKey("non-zero-key"))
	})
}
