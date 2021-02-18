package tracing

import (
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
)

func NewJaeger(serviceName string, registry prometheus.Registerer) (opentracing.Tracer, io.Closer, error) {
	// mtr := jaegerMetrics.New(jaegerMetrics.WithRegistry(registry))
	return nil, nil, nil
}
