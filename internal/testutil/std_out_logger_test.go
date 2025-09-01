package testutil_test

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"

	"go.uber.org/zap"

	"go.robertomontagna.dev/zapfluent/internal/testutil"

	. "github.com/onsi/gomega"
)

func TestStdoutLoggerForTest(t *testing.T) {
	type testCase struct {
		name           string
		options        []zap.Option
		logMessage     string
		expectedMsg    string
		expectedLevel  string
		expectedFields map[string]interface{}
	}

	testCases := []testCase{
		{
			name:           "without options",
			options:        []zap.Option{},
			logMessage:     "message 1",
			expectedMsg:    "message 1",
			expectedLevel:  "info",
			expectedFields: map[string]interface{}{},
		},
		{
			name: "with options",
			options: []zap.Option{
				zap.Fields(
					zap.String("component", "testing"),
					zap.Int("number", 123),
				),
			},
			logMessage:    "message 2",
			expectedMsg:   "message 2",
			expectedLevel: "info",
			expectedFields: map[string]interface{}{
				"component": "testing",
				"number":    float64(123), // JSON unmarshals numbers to float64
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			originalStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			logger := testutil.StdoutLoggerForTest(tc.options...)
			logger.Info(tc.logMessage)
			w.Close()

			os.Stdout = originalStdout
			var buf bytes.Buffer
			_, err := io.Copy(&buf, r)
			g.Expect(err).ToNot(HaveOccurred())

			var logOutput map[string]interface{}
			err = json.Unmarshal(buf.Bytes(), &logOutput)
			g.Expect(err).ToNot(HaveOccurred(), "should be a valid JSON")

			g.Expect(logOutput).To(HaveKeyWithValue("level", tc.expectedLevel))
			g.Expect(logOutput).To(HaveKeyWithValue("msg", tc.expectedMsg))
			for key, value := range tc.expectedFields {
				g.Expect(logOutput).To(HaveKeyWithValue(key, value))
			}
		})
	}
}
