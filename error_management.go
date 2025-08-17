package zapfluent

type ErrorManagementMode int

const (
	ErrorManagementModeUnknown ErrorManagementMode = iota
	ErrorManagementModeEarlyFailing
	ErrorManagementModeContinueDiscard
	ErrorManagementModeContinuePrintFailures
	ErrorManagementModeNeverFails
)

func (m ErrorManagementMode) String() string {
	if m < ErrorManagementModeUnknown || m > ErrorManagementModeNeverFails {
		return "Unknown"
	}
	return [...]string{
		"Unknown",
		"EarlyFailing",
		"ContinueDiscard",
		"ContinuePrintFailures",
		"NeverFails",
	}[m]
}

func IntToErrorManagementMode(value int) ErrorManagementMode {
	if value < int(ErrorManagementModeUnknown) || value > int(ErrorManagementModeNeverFails) {
		return ErrorManagementModeUnknown
	}
	return ErrorManagementMode(value)
}
