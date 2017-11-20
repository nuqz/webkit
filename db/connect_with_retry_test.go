package db

import (
	"errors"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
)

func Test__ConnectWithRetry(t *testing.T) {
	i, n, c := time.Microsecond, int8(5), int8(0)
	tconn, terr := new(sqlx.DB), errors.New("test")
	cfn := func() (*sqlx.DB, error) {
		c++
		return tconn, terr
	}

	conn, err := ConnectWithRetry(i, n, cfn)
	if c != n {
		t.Errorf("%d/%d connection attempts", n, c)
	}

	if conn != tconn {
		t.Error("should return connection returned from connecttion function")
	}

	if err != terr {
		t.Error("should return error from connection function")
	}

	c = 0
	cfn = func() (*sqlx.DB, error) {
		c++
		return tconn, nil
	}
	conn, err = ConnectWithRetry(i, n, cfn)

	if c != 1 {
		t.Error("should not retry if connection was successfull")
	}
}
