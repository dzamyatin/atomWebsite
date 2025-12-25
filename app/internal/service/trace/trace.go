package servicetrace

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

type Trace struct {
	tracer *opentracing.Tracer
}

func NewTrace(
	tracer *opentracing.Tracer,
) *Trace {
	return &Trace{
		tracer: tracer,
	}
}

type Tag struct {
	Key   string
	Value string
}

type TraceFunc func(ctx context.Context)

func (r *Trace) Trace(
	ctx context.Context,
	operation string,
	f TraceFunc,
	tags ...Tag,
) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		*r.tracer,
		operation,
	)

	for _, tag := range tags {
		span.SetTag(tag.Key, tag.Value)
	}

	f(ctx)

	span.Finish()
}
