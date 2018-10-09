package gigasecond

import "time"

// AddGigasecond will return a time object
// of the moment someone has lived for 10^9 seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(1000000000) * time.Second)
}
