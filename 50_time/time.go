package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// current time
	now := time.Now()
	p(now)

	// You can build a time struct by providing the year, month, day, etc. Times are always
	// associated with a Location, i.e. time zone.
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// You can extract the various components of the time value as expected.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// The Monday-Sunday Weekday is also available.
	p(then.Weekday())

	// These methods compare two times, testing if the first occurs before, after, or at the same
	// time as the second, respectively.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// The Sub methods returns a Duration representing the interval between two times.
	diff := now.Sub(then)
	p(diff)

	// We can compute the length of the duration in various units.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// You can use Add to advance a time by a given duration, or with a - to move backwards by a
	// duration.
	p(then.Add(diff))
	p(then.Add(-diff))
}

/*
   Test:
   ----
   $ go run 50_time/time.go
   2021-01-19 20:59:45.785784643 +0300 EAT m=+0.000042406
   2009-11-17 20:34:58.651387237 +0000 UTC
   2009
   November
   17
   20
   34
   58
   651387237
   UTC
   Tuesday
   true
   false
   false
   97941h24m47.134397406s
   97941.41309288816
   5.87648478557329e+06
   3.525890871343974e+08
   352589087134397406
   2021-01-19 17:59:45.785784643 +0000 UTC
   1998-09-15 23:10:11.516989831 +0000 UTC
*/
