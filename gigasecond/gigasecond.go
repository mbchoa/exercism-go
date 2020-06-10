// Package gigasecond provides the date and time plus 1 Gigasecond (1e9 seconds) in the future
package gigasecond

import "time"

const GIGASECOND = time.Duration(1e9) * time.Second

// AddGigasecond adds a single gigasecond to the input time and returns it
func AddGigasecond(t time.Time) time.Time {
	return t.Add(GIGASECOND)
}
