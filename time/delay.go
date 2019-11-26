package time

import (
	"time"
)

// Delay is delay duration by time
func Delay(t time.Time) time.Duration {
	return time.Now().Sub(t)
}

// DelayByInt64 is delay duration by int64
func DelayByInt64(i int64) time.Duration {
	t := time.Unix(i, 0)
	return time.Now().Sub(t)
}
