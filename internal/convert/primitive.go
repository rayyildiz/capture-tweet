package convert

import (
	"time"
)

func String(s string) *string {
	if len(s) == 0 {
		return nil
	}
	return &s
}

func Int(n int) *int {
	return &n
}

func Time(t time.Time) *time.Time {
	return &t
}
