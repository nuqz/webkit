package support

import (
	"errors"
	"testing"
	"time"
)

func Test__RetryOnErr(t *testing.T) {
	i, n, c := time.Microsecond, 5, 0
	tr, terr := int64(1), errors.New("test")
	fn := func() (interface{}, error) {
		c++
		return tr, terr
	}

	r, err := RetryOnErr(i, n, fn)
	if c != n {
		t.Errorf("%d/%d connection attempts", n, c)
	}

	if r.(int64) != tr {
		t.Error("should return value returned from function")
	}

	if err != terr {
		t.Error("should return error returned from function")
	}

	c = 0
	fn = func() (interface{}, error) {
		c++
		return tr, nil
	}
	r, err = RetryOnErr(i, n, fn)

	if c != 1 {
		t.Error("should not retry if function didn't return error")
	}
}
