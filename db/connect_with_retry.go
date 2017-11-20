package db

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// ConnectWithRetry tries to connect to the database using cfn function n times
// with interval i. Returns database connection after first successful attempt.
// Returns the last connection error if all n tries was unsuccessfull.
func ConnectWithRetry(
	i time.Duration,
	n int8,
	cfn func() (*sqlx.DB, error),
) (dbConn *sqlx.DB, err error) {
	try := int8(1)
	for range time.Tick(i) {
		if try > n {
			break
		}

		if dbConn, err = cfn(); err == nil {
			break
		}

		try++
	}
	return dbConn, err
}
