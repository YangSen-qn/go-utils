package times

import (
	"time"
)

// FormatTime formats a time.Time object to a string with the specified layout
func FormatTime(t time.Time, layout string) string {
	return t.Format(layout)
}

// ParseTime parses a time string according to the specified layout
func ParseTime(s, layout string) (time.Time, error) {
	return time.Parse(layout, s)
}

// GetCurrentTime returns the current time in the specified location
func GetCurrentTime(loc *time.Location) time.Time {
	return time.Now().In(loc)
}

// CalculateDuration calculates the duration between two time.Time objects
func CalculateDuration(start, end time.Time) time.Duration {
	return end.Sub(start)
}

// AddDuration adds a duration to a time.Time object
func AddDuration(t time.Time, d time.Duration) time.Time {
	return t.Add(d)
}

// IsFuture checks if a time is in the future
func IsFuture(t time.Time) bool {
	return t.After(time.Now())
}

// IsPast checks if a time is in the past
func IsPast(t time.Time) bool {
	return t.Before(time.Now())
}

// FormatUnixTimestamp formats a Unix timestamp to a time.Time object
func FormatUnixTimestamp(ts int64, layout string) string {
	return time.Unix(ts, 0).Format(layout)
}

// ParseUnixTimestamp parses a Unix timestamp from a string
func ParseUnixTimestamp(s string) (int64, error) {
	t, err := time.Parse(time.UnixDate, s)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
