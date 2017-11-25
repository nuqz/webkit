package support

import "time"

// RetryOnErr calls fn function no more than n times with interval i after each
// attempt and returns the values that were returned when the fn function was
// last called.
func RetryOnErr(
	i time.Duration,
	n int,
	fn func() (interface{}, error),
) (r interface{}, err error) {
	try := 1
	for range time.Tick(i) {
		if try > n {
			break
		}

		if r, err = fn(); err == nil {
			break
		}

		try++
	}
	return r, err
}
