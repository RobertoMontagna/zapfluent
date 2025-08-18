package zapfluent

import (
	"go.robertomontagna.dev/zapfluent/config"
	"go.uber.org/multierr"
)

type ErrorHandler struct {
	err  error
	mode config.ErrorHandlingMode
}

func NewErrorHandler(cfg config.ErrorHandlingConfiguration) *ErrorHandler {
	return &ErrorHandler{
		mode: cfg.Mode(),
	}
}

func (h *ErrorHandler) ShouldSkip() bool {
	return h.mode == config.ErrorHandlingModeEarlyFailing && h.err != nil
}

func (h *ErrorHandler) ManageError(newErr error) {
	if newErr == nil {
		return
	}
	if h.mode == config.ErrorHandlingModeEarlyFailing && h.err != nil {
		return
	}
	h.err = multierr.Append(h.err, newErr)
}

func (h *ErrorHandler) AggregatedError() error {
	return h.err
}
