package opentelemetry

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type OpenTelemetry struct {
	ServiceName  string
	CollectorURL string
}

func NewOpenTelemetry(serviceName, collectorURL string) *OpenTelemetry {
	return &OpenTelemetry{
		ServiceName:  serviceName,
		CollectorURL: collectorURL,
	}
}

func (op *OpenTelemetry) InitProvider() (func(context.Context) error, error) {
	ctx := context.Background()
	res, err := op.NewResource(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	traceExporter, err := op.NewExporter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	traceProvider := op.NewTraceProvider(traceExporter, res)

	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return traceProvider.Shutdown, nil
}

func (op *OpenTelemetry) NewResource(ctx context.Context) (*resource.Resource, error) {
	return resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(op.ServiceName),
		),
	)
}

func (op *OpenTelemetry) NewExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	fmt.Println(op.CollectorURL)
	conn, err := grpc.DialContext(ctx, op.CollectorURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to collector: %w", err)
	}

	return otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
}

func (op *OpenTelemetry) NewTraceProvider(traceExporter *otlptrace.Exporter, res *resource.Resource) *sdktrace.TracerProvider {
	batcher := sdktrace.NewBatchSpanProcessor(traceExporter)

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(batcher),
	)

	return traceProvider
}
