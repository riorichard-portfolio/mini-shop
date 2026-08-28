package entity

import (
	"errors"
	"time"
)

type Order struct {
	id          string
	customerID  string
	productID   string
	productName string
	quantity    int
	createdAt   time.Time
	status      string
}

func NewOrder(
	id string,
	customerID string,
	productID string,
	productName string,
	quantity int,
	createdAt time.Time,
	status string,
) (*Order, error) {
	if quantity < 1 {
		return nil, errors.New("INVALID_QUANTITY")
	}
	if status != "PENDING" && status != "CANCELLED" && status != "COMPLETE" {
		return nil, errors.New("INVALID_ORDER_STATUS")
	}
	return &Order{
		id:          id,
		customerID:  customerID,
		productID:   productID,
		productName: productName,
		quantity:    quantity,
		createdAt:   createdAt,
		status:      status,
	}, nil
}

func (o *Order) ID() string           { return o.id }
func (o *Order) CustomerID() string   { return o.customerID }
func (o *Order) ProductID() string    { return o.productID }
func (o *Order) ProductName() string  { return o.productName }
func (o *Order) Quantity() int        { return o.quantity }
func (o *Order) CreatedAt() time.Time { return o.createdAt }
func (o *Order) Status() string       { return o.status }

func (o *Order) Cancel() {
	o.status = "CANCELLED"
}

func (o *Order) Complete() {
	o.status = "COMPLETE"
}
