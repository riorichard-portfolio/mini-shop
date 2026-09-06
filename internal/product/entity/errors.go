package entity

import "errors"

var (
	InvalidQuantityErr   = errors.New("INVALID_QUANTITY")
	UnauthorizedSellerErr = errors.New("UNAUTHORIZED_SELLER")
)
