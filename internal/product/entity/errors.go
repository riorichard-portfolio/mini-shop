package entity

import "errors"

var (
	InvalidStockErr      = errors.New("INVALID_STOCK")
	InvalidQuantityErr   = errors.New("INVALID_QUANTITY")
	InsufficientStockErr = errors.New("INSUFFICIENT_STOCK")
)
