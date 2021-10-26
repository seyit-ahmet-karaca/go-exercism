// Package gigasecond should have a package comment that summarizes what it's about.
package gigasecond

import "time"

const aGigaSecond = 1000000000

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * aGigaSecond)
}
