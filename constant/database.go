package constant

import "time"

// all the databases share this settings
const (
	DatabaseHost     = "0.0.0.0"
	DatabaseName     = "cdc-observer"
	DatabaseUsername = "root"
	DatabasePassword = "cdc-observer-password"
)

// retry times and interval for the database connection
const (
	RetryTimes    = 10
	RetryInterval = 1 * time.Second
)
