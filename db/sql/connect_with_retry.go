package sql

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nuqz/webkit/support"
)

// ConnectWithRetry tries to connect to the database using cfn function no more
// than times with interval i after each attempt. Returns database connection
// after first successful attempt. Returns the last connection error when all n
// tries were unsuccessful. Panics, when the value that was returned from cfn
// function is not *sqlx.DB.
func ConnectWithRetry(
	i time.Duration,
	n int,
	cfn func() (interface{}, error),
) (interface{}, error) {
	r, err := support.RetryOnErr(i, n, cfn)
	return r.(*sqlx.DB), err
}
