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

func NewTag(name string, value string) Tag {
	return Tag{
		Key:   name,
		Value: value,
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
	var options []opentracing.StartSpanOption
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan != nil {
		options = append(options, opentracing.ChildOf(parentSpan.Context()))
	}

	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		*r.tracer,
		operation,
		options...,
	)

	for _, tag := range tags {
		span.SetTag(tag.Key, tag.Value)
	}

	f(ctx)

	span.Finish()
}
