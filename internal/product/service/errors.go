package service

import "errors"

var(
	InsufficientStockErr = errors.New("INSUFFICIENT_STOCK")
	InvalidProductErr = errors.New("INVALID_PRODUCT")
)