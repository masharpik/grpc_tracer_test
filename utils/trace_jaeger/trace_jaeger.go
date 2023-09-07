package tracejaeger

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func NewProvider(endpoint string, serviceName string) (prv *trace.TracerProvider, err error) {
	exp, err := jaeger.New(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)),
	)
	if err != nil {
		return
	}

	prv = trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)

	otel.SetTracerProvider(prv)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return
}

func Close(prv *trace.TracerProvider, ctx context.Context) error {
	return prv.Shutdown(ctx)
}
