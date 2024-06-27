package interfaces

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type Tracer interface {
	StartSpan(context.Context, string, ...trace.SpanStartOption) (trace.Span, context.Context)
}
