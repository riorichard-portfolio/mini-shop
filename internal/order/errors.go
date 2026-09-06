package order

import "errors"

var(
	InconsistentStatusChangesErr = errors.New("INCONSISTENT_STATUS_CHANGES")
)