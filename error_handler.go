package zapfluent

import (
	"go.robertomontagna.dev/zapfluent/config"
	"go.robertomontagna.dev/zapfluent/fluentfield"
	"go.robertomontagna.dev/zapfluent/functional/optional"
	"go.uber.org/multierr"
)

// errorHandler is responsible for managing errors that occur during field
// encoding. It supports different strategies, such as aggregating multiple
// errors or failing on the first error. It can also create fallback fields
// when an encoding operation fails.
type errorHandler struct {
	err error
	cfg config.ErrorHandlingConfiguration
}

// newErrorHandler creates a new errorHandler with the given configuration.
func newErrorHandler(cfg config.ErrorHandlingConfiguration) *errorHandler {
	return &errorHandler{
		cfg: cfg,
	}
}

// shouldSkip determines whether the fluent chain should stop processing new
// fields. This is used in "EarlyFailing" mode, where processing halts after
// the first error.
func (eh *errorHandler) shouldSkip() bool {
	return eh.cfg.Mode() == config.ErrorHandlingModeEarlyFailing && eh.err != nil
}

// process handles an error that occurred while encoding a field.
//
// If the error is non-nil, it is aggregated. If a fallback factory is
// configured, it creates and returns a fallback field. Otherwise, it returns
// an empty Optional.
func (eh *errorHandler) process(field fluentfield.Field, err error) optional.Optional[fluentfield.Field] {
	if err == nil {
		return optional.Empty[fluentfield.Field]()
	}

	eh.aggregateError(err)

	return optional.Map(
		eh.cfg.FallbackFactory(),
		func(factory config.FallbackFieldFactory) fluentfield.Field {
			return factory(field.Name(), err)
		},
	)
}

// aggregateError appends a new error to the collection of errors.
// It uses `multierr.Append` to create or add to an error chain.
func (eh *errorHandler) aggregateError(newErr error) {
	if newErr == nil {
		return
	}
	eh.err = multierr.Append(eh.err, newErr)
}

// aggregatedError returns the final error, which may be an aggregation of
// multiple errors collected during processing. It returns nil if no errors
// were encountered.
func (eh *errorHandler) aggregatedError() error {
	return eh.err
}
