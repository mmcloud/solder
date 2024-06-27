package tracer

import (
	"context"

	"github.com/mmcloud/solder/pkg/interfaces"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
)

const (
	name    = "instrumentation/package/name"
	version = "0.1.0"
)

type contextKey string

const spanKey contextKey = "span"

// Instrumentation is an optional struct if you don't need specific instrumentations.
type Instrumentation struct {
	tracer trace.Tracer
}

func NewInstrumentation(tp trace.TracerProvider) *Instrumentation {
	if tp == nil {
		// Create a default TracerProvider if none is provided
		tp = sdktrace.NewTracerProvider(
			sdktrace.WithResource(resource.Default()), // Use the default resource for now
		)
	}
	return &Instrumentation{
		tracer: tp.Tracer(name, trace.WithInstrumentationVersion(version)),
	}
}

// Tracer provides a wrapper around the OpenTelemetry TracerProvider
type Tracer struct {
	embedded.TracerProvider
}

func NewTracer(tracerProvider trace.TracerProvider) *Tracer {
	return &Tracer{TracerProvider: tracerProvider}
}

var _ interfaces.Tracer = (*Tracer)(nil) // Ensure Tracer implements the interface

// StartSpan starts a new span with the given name and options.
func (t *Tracer) StartSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) (trace.Span, context.Context) {
	opts = append(opts, trace.WithSpanKind(trace.SpanKindInternal))

	ctx, span := tracer.Start(ctx, name, opts...)
	ctx = context.WithValue(ctx, spanKey, span)
	return span, ctx
}
