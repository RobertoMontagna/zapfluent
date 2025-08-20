package config

import "go.robertomontagna.dev/zapfluent/util/enum_util"

// ErrorHandlingMode defines the strategy for handling errors that occur during
// field encoding.
type ErrorHandlingMode int

const (
	// ErrorHandlingModeUnknown is the default, zero-value handling mode.
	ErrorHandlingModeUnknown ErrorHandlingMode = iota
	// ErrorHandlingModeEarlyFailing stops processing fields as soon as the
	// first error is encountered. The final aggregated error will contain
	// only the first error that occurred.
	ErrorHandlingModeEarlyFailing
	// ErrorHandlingModeContinue continues processing fields even after errors
	// occur. All errors are collected and returned as a single aggregated
	// error. This is the default behavior.
	ErrorHandlingModeContinue
)

const (
	ErrorHandlingModeUnknownString      = "Unknown"
	ErrorHandlingModeEarlyFailingString = "EarlyFailing"
	ErrorHandlingModeContinueString     = "Continue"
)

var errorHandlingModeEnum = enum_util.NewUtilEnum(
	map[ErrorHandlingMode]string{
		ErrorHandlingModeUnknown:      ErrorHandlingModeUnknownString,
		ErrorHandlingModeEarlyFailing: ErrorHandlingModeEarlyFailingString,
		ErrorHandlingModeContinue:     ErrorHandlingModeContinueString,
	},
	ErrorHandlingModeUnknown,
)

// String returns the string representation of the ErrorHandlingMode.
func (m ErrorHandlingMode) String() string {
	return errorHandlingModeEnum.String(m)
}

// IntToErrorHandlingMode converts an integer to an ErrorHandlingMode.
// If the integer does not correspond to a valid mode, it returns
// ErrorHandlingModeUnknown.
func IntToErrorHandlingMode(value int) ErrorHandlingMode {
	return errorHandlingModeEnum.FromInt(value)
}
