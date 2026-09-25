package usecase

import "mini-shop/internal/pkg/err/bizerr"

var(
	InconsistentStatusChangesErr = bizerr.BadRequestErr("INCONSISTENT_STATUS_CHANGES")
)