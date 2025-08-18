package zapfluent

import (
	"go.robertomontagna.dev/zapfluent/config"
	"go.uber.org/multierr"
)

type ErrorHandler struct {
	err error
	cfg config.ErrorHandlingConfiguration
}

func NewErrorHandler(cfg config.ErrorHandlingConfiguration) *ErrorHandler {
	return &ErrorHandler{
		cfg: cfg,
	}
}

func (h *ErrorHandler) ShouldSkip() bool {
	return h.cfg.Mode() == config.ErrorHandlingModeEarlyFailing && h.err != nil
}

func (h *ErrorHandler) AggregateError(newErr error) {
	if newErr == nil {
		return
	}
	h.err = multierr.Append(h.err, newErr)
}

func (h *ErrorHandler) AggregatedError() error {
	return h.err
}
