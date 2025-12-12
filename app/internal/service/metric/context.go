package metric

import "context"

var LabelKey = &struct{}{}

func MetricContext(ctx context.Context, label string) context.Context {
	return context.WithValue(ctx, LabelKey, label)
}

func GetMetricContextLabel(ctx context.Context) string {
	v, ok := ctx.Value(LabelKey).(string)
	if ok {
		return v
	}
	
	return ""
}
