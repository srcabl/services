package tracing

import (
	"io"

	opentracing "github.com/opentracing/opentracing-go"
)

type Tracer struct {
	Tracer opentracing.Tracer
	Closer io.Closer
}
