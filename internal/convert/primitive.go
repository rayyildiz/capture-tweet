package convert

import (
	"time"
)

func String(s string) *string {
	return &s
}

func Int(n int) *int {
	return &n
}

func Time(t time.Time) *time.Time {
	return &t
}
