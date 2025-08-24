package core

import (
	"go.uber.org/multierr"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/internal/functional/optional"
)

type ErrorHandler struct {
	cfg        *ErrorHandlingConfiguration
	enc        zapcore.ObjectEncoder
	totalError error
}

func NewErrorHandler(
	cfg *ErrorHandlingConfiguration,
	enc zapcore.ObjectEncoder,
) *ErrorHandler {
	return &ErrorHandler{
		cfg: cfg,
		enc: enc,
	}
}

func (h *ErrorHandler) ShouldSkip() bool {
	return h.cfg.Mode() == ErrorHandlingModeEarlyFailing && h.totalError != nil
}

func (h *ErrorHandler) handleError(field Field, err error) optional.Optional[Field] {
	if err == nil {
		return optional.Empty[Field]()
	}

	h.aggregateError(err)

	return optional.Map(
		h.cfg.FallbackFactory(),
		func(factory FallbackFieldFactory) Field {
			return factory(field.Name(), err)
		},
	)
}

type FieldEncodingErrorManager func()

func (h *ErrorHandler) EncodeField(field Field) FieldEncodingErrorManager {
	if h.ShouldSkip() {
		return func() {
			// This function is intentionally left empty.
			// When the ErrorHandler is in EarlyFailing mode and an error has already occurred,
			// subsequent field encoding operations should be skipped entirely.
			// Returning a no-op function is the most efficient way to achieve this.
		}
	}
	maybeFallbackField := h.handleError(field, field.Encode(h.enc))

	return func() {
		maybeEncodingError := optional.FlatMap(maybeFallbackField, h.encodeAndLift)
		maybeFallbackFailed := optional.Map(maybeEncodingError, func(_ error) Field {
			return String(field.Name(), h.cfg.FallbackErrorMessage)
		})
		optional.FlatMap(maybeFallbackFailed, h.encodeAndLift)
	}
}

func (h *ErrorHandler) encodeAndLift(field Field) optional.Optional[error] {
	err := field.Encode(h.enc)
	h.aggregateError(err)
	return optional.OfError(err)
}

func (h *ErrorHandler) aggregateError(newErr error) {
	h.totalError = multierr.Append(h.totalError, newErr)
}

func (h *ErrorHandler) AggregatedError() error {
	return h.totalError
}
