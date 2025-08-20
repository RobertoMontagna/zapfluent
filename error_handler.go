package zapfluent

import (
	"go.robertomontagna.dev/zapfluent/config"
	"go.robertomontagna.dev/zapfluent/fluentfield"
	"go.robertomontagna.dev/zapfluent/functional/optional"
	"go.uber.org/multierr"
)

// ErrorHandler is responsible for managing errors that occur during field
// encoding. It supports different strategies, such as aggregating multiple
// errors or failing on the first error. It can also create fallback fields
// when an encoding operation fails.
type ErrorHandler struct {
	err error
	cfg config.ErrorHandlingConfiguration
}

// NewErrorHandler creates a new ErrorHandler with the given configuration.
func NewErrorHandler(cfg config.ErrorHandlingConfiguration) *ErrorHandler {
	return &ErrorHandler{
		cfg: cfg,
	}
}

// ShouldSkip determines whether the fluent chain should stop processing new
// fields. This is used in "EarlyFailing" mode, where processing halts after
// the first error.
func (eh *ErrorHandler) ShouldSkip() bool {
	return eh.cfg.Mode() == config.ErrorHandlingModeEarlyFailing && eh.err != nil
}

// Process handles an error that occurred while encoding a field.
//
// If the error is non-nil, it is aggregated. If a fallback factory is
// configured, it creates and returns a fallback field. Otherwise, it returns
// an empty Optional.
func (eh *ErrorHandler) Process(field fluentfield.Field, err error) optional.Optional[fluentfield.Field] {
	if err == nil {
		return optional.Empty[fluentfield.Field]()
	}

	eh.AggregateError(err)

	return optional.Map(
		eh.cfg.FallbackFactory(),
		func(factory config.FallbackFieldFactory) fluentfield.Field {
			return factory(field.Name(), err)
		},
	)
}

// AggregateError appends a new error to the collection of errors.
// It uses `multierr.Append` to create or add to an error chain.
func (eh *ErrorHandler) AggregateError(newErr error) {
	if newErr == nil {
		return
	}
	eh.err = multierr.Append(eh.err, newErr)
}

// AggregatedError returns the final error, which may be an aggregation of
// multiple errors collected during processing. It returns nil if no errors
// were encountered.
func (eh *ErrorHandler) AggregatedError() error {
	return eh.err
}
