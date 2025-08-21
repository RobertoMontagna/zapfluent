package zapfluent

import (
	"go.uber.org/multierr"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/config"
	"go.robertomontagna.dev/zapfluent/fluentfield"
	"go.robertomontagna.dev/zapfluent/functional/optional"
)

type errorHandler struct {
	cfg        config.ErrorHandlingConfiguration
	enc        zapcore.ObjectEncoder
	totalError error
}

func newErrorHandler(
	cfg config.ErrorHandlingConfiguration,
	enc zapcore.ObjectEncoder,
) *errorHandler {
	return &errorHandler{
		cfg: cfg,
		enc: enc,
	}
}

func (h *errorHandler) shouldSkip() bool {
	return h.cfg.Mode() == config.ErrorHandlingModeEarlyFailing && h.totalError != nil
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

type fieldEncodingErrorManager func()

func (h *errorHandler) encodeField(field fluentfield.Field) fieldEncodingErrorManager {
	if h.shouldSkip() {
		return func() {}
	}
	maybeFallbackField := h.handleError(field, field.Encode(h.enc))

	return func() {
		maybeEncodingError := optional.FlatMap(maybeFallbackField, h.encodeAndLift)
		maybeFallbackFailed := optional.Map(maybeEncodingError, func(_ error) fluentfield.Field {
			return fluentfield.String(field.Name(), h.cfg.FallbackErrorMessage)
		})
		optional.FlatMap(maybeFallbackFailed, h.encodeAndLift)
	}
}

func (h *errorHandler) encodeAndLift(field fluentfield.Field) optional.Optional[error] {
	err := field.Encode(h.enc)
	h.aggregateError(err)
	return optional.OfError(err)
}

func (h *errorHandler) aggregateError(newErr error) {
	h.totalError = multierr.Append(h.totalError, newErr)
}

func (h *errorHandler) aggregatedError() error {
	return h.totalError
}
