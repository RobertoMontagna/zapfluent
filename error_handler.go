package zapfluent

import (
	"go.uber.org/multierr"

	"go.robertomontagna.dev/zapfluent/config"
	"go.robertomontagna.dev/zapfluent/fluentfield"
	"go.robertomontagna.dev/zapfluent/functional/optional"
)

type errorHandler struct {
	err error
	cfg config.ErrorHandlingConfiguration
}

func newErrorHandler(cfg config.ErrorHandlingConfiguration) *errorHandler {
	return &errorHandler{
		cfg: cfg,
	}
}

func (h *errorHandler) shouldSkip() bool {
	return h.cfg.Mode() == config.ErrorHandlingModeEarlyFailing && h.err != nil
}

func (h *errorHandler) handleError(field fluentfield.Field, err error) optional.Optional[fluentfield.Field] {
	if err == nil {
		return optional.Empty[fluentfield.Field]()
	}

	h.aggregateError(err)

	return optional.Map(
		h.cfg.FallbackFactory(),
		func(factory config.FallbackFieldFactory) fluentfield.Field {
			return factory(field.Name(), err)
		},
	)
}

func (h *errorHandler) aggregateError(newErr error) {
	if newErr == nil {
		return
	}
	h.err = multierr.Append(h.err, newErr)
}

func (h *errorHandler) aggregatedError() error {
	return h.err
}
