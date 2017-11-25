package sql

import (
	"errors"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
)

func Test__ConnectWithRetry(t *testing.T) {
	tconn, terr := new(sqlx.DB), errors.New("test error")
	fn := func() (interface{}, error) {
		return tconn, terr
	}

	r, err := ConnectWithRetry(time.Microsecond, 5, fn)
	if err != terr {
		t.Errorf("expected error %s, but got %s", terr.Error(), err.Error())
	}

	if conn, ok := r.(*sqlx.DB); !ok || conn != tconn {
		t.Errorf("expected db connection %p, but got %p", tconn, conn)
	}

	p := false
	defer func() {
		if err := recover(); err != nil {
			p = true
		}

		if !p {
			t.Error("should panic, when database connection is not *sqlx.DB")
		}
	}()

	fn = func() (interface{}, error) {
		return nil, nil
	}
	_, _ = ConnectWithRetry(time.Microsecond, 5, fn)
}
