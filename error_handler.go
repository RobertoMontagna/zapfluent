package zapfluent

import (
	"go.uber.org/multierr"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/internal/functional/optional"
	"go.robertomontagna.dev/zapfluent/pkg/core"
)

type errorHandler struct {
	cfg        *core.ErrorHandlingConfiguration
	enc        zapcore.ObjectEncoder
	totalError error
}

func newErrorHandler(
	cfg *core.ErrorHandlingConfiguration,
	enc zapcore.ObjectEncoder,
) *errorHandler {
	return &errorHandler{
		cfg: cfg,
		enc: enc,
	}
}

func (h *errorHandler) shouldSkip() bool {
	return h.cfg.Mode() == core.ErrorHandlingModeEarlyFailing && h.totalError != nil
}

func (h *errorHandler) handleError(field core.Field, err error) optional.Optional[core.Field] {
	if err == nil {
		return optional.Empty[core.Field]()
	}

	h.aggregateError(err)

	return optional.Map(
		h.cfg.FallbackFactory(),
		func(factory core.FallbackFieldFactory) core.Field {
			return factory(field.Name(), err)
		},
	)
}

type fieldEncodingErrorManager func()

func (h *errorHandler) encodeField(field core.Field) fieldEncodingErrorManager {
	if h.shouldSkip() {
		return func() {}
	}
	maybeFallbackField := h.handleError(field, field.Encode(h.enc))

	return func() {
		maybeEncodingError := optional.FlatMap(maybeFallbackField, h.encodeAndLift)
		maybeFallbackFailed := optional.Map(maybeEncodingError, func(_ error) core.Field {
			return core.String(field.Name(), h.cfg.FallbackErrorMessage)
		})
		optional.FlatMap(maybeFallbackFailed, h.encodeAndLift)
	}
}

func (h *errorHandler) encodeAndLift(field core.Field) optional.Optional[error] {
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
