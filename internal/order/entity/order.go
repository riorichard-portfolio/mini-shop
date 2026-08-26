package entity

import (
	"time"
)

type Order struct {
	id          string
	customerID  string
	productID   string
	productName string
	quantity    int
	createdAt   time.Time
}

func NewOrder(
	id string,
	customerID string,
	productID string,
	productName string,
	quantity int,
	createdAt time.Time,
) *Order {
	return &Order{
		id:          id,
		customerID:  customerID,
		productID:   productID,
		productName: productName,
		quantity:    quantity,
		createdAt:   createdAt,
	}
}

func (o *Order) ID() string           { return o.id }
func (o *Order) CustomerID() string   { return o.customerID }
func (o *Order) ProductID() string    { return o.productID }
func (o *Order) ProductName() string  { return o.productName }
func (o *Order) Quantity() int        { return o.quantity }
func (o *Order) CreatedAt() time.Time { return o.createdAt }
