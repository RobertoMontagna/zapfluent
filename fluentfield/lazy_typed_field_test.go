package fluentfield_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent/fluentfield"
	"go.uber.org/zap/zapcore"
)

const (
	testFieldName = "test-field"
)

func TestLazyTypedField_Encode(t *testing.T) {
	t.Run("when value is present, it encodes the value", func(t *testing.T) {
		functions := fluentfield.TypeFieldFunctions[string]{
			EncodeFunc: func(enc zapcore.ObjectEncoder, name string, value string) error {
				enc.AddString(name, value)
				return nil
			},
			IsNonZero: func(s string) bool { return s != "" },
		}
		field := fluentfield.NewTypedField(functions, testFieldName, "test-value")
		enc := zapcore.NewMapObjectEncoder()

		err := field.Encode(enc)

		assert.NoError(t, err)
		assert.Equal(t, "test-value", enc.Fields[testFieldName])
	})

	t.Run("when value is not present, it does not encode anything", func(t *testing.T) {
		functions := fluentfield.TypeFieldFunctions[string]{
			EncodeFunc: func(enc zapcore.ObjectEncoder, name string, value string) error {
				enc.AddString(name, value)
				return nil
			},
			IsNonZero: func(s string) bool { return s != "" },
		}
		field := fluentfield.NewTypedField(functions, testFieldName, "test-value").
			Filter(func(s string) bool { return false }) // This will make the value not present
		enc := zapcore.NewMapObjectEncoder()

		err := field.Encode(enc)

		assert.NoError(t, err)
		assert.Empty(t, enc.Fields)
	})
}

func TestLazyTypedField_Name(t *testing.T) {
	functions := fluentfield.TypeFieldFunctions[string]{}
	field := fluentfield.NewTypedField(functions, testFieldName, "test-value")

	assert.Equal(t, testFieldName, field.Name())
}

func TestLazyTypedField_Filter(t *testing.T) {
	functions := fluentfield.TypeFieldFunctions[string]{
		EncodeFunc: func(enc zapcore.ObjectEncoder, name string, value string) error {
			enc.AddString(name, value)
			return nil
		},
		IsNonZero: func(s string) bool { return s != "" },
	}
	field := fluentfield.NewTypedField(functions, testFieldName, "test-value")

	t.Run("when condition is met, it keeps the value", func(t *testing.T) {
		filteredField := field.Filter(func(s string) bool { return true })
		enc := zapcore.NewMapObjectEncoder()
		_ = filteredField.Encode(enc)
		assert.NotEmpty(t, enc.Fields)
	})

	t.Run("when condition is not met, it removes the value", func(t *testing.T) {
		filteredField := field.Filter(func(s string) bool { return false })
		enc := zapcore.NewMapObjectEncoder()
		_ = filteredField.Encode(enc)
		assert.Empty(t, enc.Fields)
	})
}

func TestLazyTypedField_NonZero(t *testing.T) {
	functions := fluentfield.TypeFieldFunctions[string]{
		EncodeFunc: func(enc zapcore.ObjectEncoder, name string, value string) error {
			enc.AddString(name, value)
			return nil
		},
		IsNonZero: func(s string) bool { return s != "" },
	}

	t.Run("when value is not zero, it keeps the value", func(t *testing.T) {
		field := fluentfield.NewTypedField(functions, testFieldName, "test-value").NonZero()
		enc := zapcore.NewMapObjectEncoder()
		_ = field.Encode(enc)
		assert.NotEmpty(t, enc.Fields)
	})

	t.Run("when value is zero, it removes the value", func(t *testing.T) {
		field := fluentfield.NewTypedField(functions, testFieldName, "").NonZero()
		enc := zapcore.NewMapObjectEncoder()
		_ = field.Encode(enc)
		assert.Empty(t, enc.Fields)
	})
}

func TestLazyTypedField_Format(t *testing.T) {
	functions := fluentfield.TypeFieldFunctions[int]{
		EncodeFunc: func(enc zapcore.ObjectEncoder, name string, value int) error {
			enc.AddInt(name, value)
			return nil
		},
	}
	field := fluentfield.NewTypedField(functions, testFieldName, 5)
	formattedField := field.Format(func(i int) string { return "formatted" })
	enc := zapcore.NewMapObjectEncoder()

	err := formattedField.Encode(enc)

	assert.NoError(t, err)
	assert.Equal(t, "formatted", enc.Fields[testFieldName])
}
