package entity

import "errors"

var (
	InvalidQuantityErr = errors.New("INVALID_QUANTITY")
	InvalidOrderStatusErr = errors.New("INVALID_ORDER_STATUS")
	UnauthorizedSellerErr = errors.New("UNAUTHORIZED_SELLER")
	NotAllowedProcessErr = errors.New("NOT_ALLOWED_PROCESS")
	StatusChangeInvalidErr = errors.New("STATUS_CHANGE_INVALID")
)