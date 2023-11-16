// package gigasecond implements a function to add a gigasecond to a specified time
package gigasecond

import "time"

// AddGigasecond adds a gigasecond to a specified time
func AddGigasecond(t time.Time) time.Time {
	duration, _ := time.ParseDuration("1000000000s")
	return t.Add(duration)
}
