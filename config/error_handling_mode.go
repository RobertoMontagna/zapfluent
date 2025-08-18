package config

import "go.robertomontagna.dev/zapfluent/util/enum_util"

type ErrorHandlingMode int

const (
	ErrorHandlingModeUnknown ErrorHandlingMode = iota
	ErrorHandlingModeEarlyFailing
	ErrorHandlingModeContinue
)

var errorHandlingModeEnum = enum_util.NewUtilEnum(
	map[ErrorHandlingMode]string{
		ErrorHandlingModeUnknown:      "Unknown",
		ErrorHandlingModeEarlyFailing: "EarlyFailing",
		ErrorHandlingModeContinue:     "Continue",
	},
	ErrorHandlingModeUnknown,
	ErrorHandlingModeUnknown,
	ErrorHandlingModeContinue,
)

func (m ErrorHandlingMode) String() string {
	return errorHandlingModeEnum.String(m)
}

func IntToErrorHandlingMode(value int) ErrorHandlingMode {
	return errorHandlingModeEnum.FromInt(value)
}
