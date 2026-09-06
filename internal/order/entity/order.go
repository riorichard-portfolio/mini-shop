package entity

import (
	"time"
)

type Order struct {
	id            string
	customerID    string
	productID     string
	sellerID      string
	productName   string
	quantity      int
	createdAt     time.Time
	currentStatus string
	statusTo      string
}

func NewOrder(
	id string,
	customerID string,
	productID string,
	sellerID string,
	productName string,
	quantity int,
	createdAt time.Time,
	currentStatus string,
) (Order, error) {
	if quantity < 1 {
		return Order{}, InvalidQuantityErr
	}
	if currentStatus != CancelledStatus &&
		currentStatus != PendingStatus &&
		currentStatus != CompleteStatus {
		return Order{}, InvalidOrderStatusErr
	}
	return Order{
		id:            id,
		customerID:    customerID,
		productID:     productID,
		sellerID:      sellerID,
		productName:   productName,
		quantity:      quantity,
		createdAt:     createdAt,
		currentStatus: currentStatus,
	}, nil
}

func (o *Order) ID() string            { return o.id }
func (o *Order) CustomerID() string    { return o.customerID }
func (o *Order) ProductID() string     { return o.productID }
func (o *Order) SellerID() string      { return o.sellerID }
func (o *Order) ProductName() string   { return o.productName }
func (o *Order) Quantity() int         { return o.quantity }
func (o *Order) CreatedAt() time.Time  { return o.createdAt }

func (o *Order) StatusTo() (string, error) {
	if o.statusTo == "" {
		return o.currentStatus, StatusChangeInvalidErr
	}
	return o.statusTo, nil
}

func (o *Order) Cancel(sellerID string) error {
	if o.currentStatus != PendingStatus {
		return NotAllowedProcessErr
	}
	if o.sellerID != sellerID {
		return UnauthorizedSellerErr
	}
	o.statusTo = CancelledStatus
	return nil
}

func (o *Order) Complete(sellerID string) error {
	if o.currentStatus != PendingStatus {
		return NotAllowedProcessErr
	}
	if o.sellerID != sellerID {
		return UnauthorizedSellerErr
	}
	o.statusTo = CompleteStatus
	return nil
}
