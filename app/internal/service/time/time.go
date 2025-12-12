package servicetime

import "time"

type ITime interface {
	Now() time.Time
}

type Time struct{}

func NewTime() *Time {
	return &Time{}
}

func (t *Time) Now() time.Time {
	return time.Now()
}
