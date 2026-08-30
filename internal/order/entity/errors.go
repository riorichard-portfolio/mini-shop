package entity

import "errors"

var (
	InvalidQuantityErr = errors.New("INVALID_QUANTITY")
	InvalidOrderStatusErr = errors.New("INVALID_ORDER_STATUS")
)