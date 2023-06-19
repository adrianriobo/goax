package delay

import "time"

const (
	SMALL  time.Duration = 1 * time.Second
	MEDIUM time.Duration = 3 * time.Second
	LONG   time.Duration = 10 * time.Second
	XLONG  time.Duration = 25 * time.Second
)

func Delay(d time.Duration) {
	time.Sleep(d)
}
