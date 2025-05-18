package main

import "time"

// wpmDuration converts WPM into tick duration
func wpmDuration(n int) time.Duration {
	return time.Minute / time.Duration(n*50)
}
