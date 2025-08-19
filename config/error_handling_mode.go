package config

import "go.robertomontagna.dev/zapfluent/util/enum_util"

type ErrorHandlingMode int

const (
	ErrorHandlingModeUnknown ErrorHandlingMode = iota
	ErrorHandlingModeEarlyFailing
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

func (m ErrorHandlingMode) String() string {
	return errorHandlingModeEnum.String(m)
}

func IntToErrorHandlingMode(value int) ErrorHandlingMode {
	return errorHandlingModeEnum.FromInt(value)
}
