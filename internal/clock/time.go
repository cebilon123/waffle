package clock

import "time"

var Now = nowFunc

func nowFunc() time.Time {
	return time.Now()
}
