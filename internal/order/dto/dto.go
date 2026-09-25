package dto

type MakeOrderInput struct {
	CustomerID string
	ProductID  string
	Quantity   int
}

type CompleteOrderInput struct {
	OrderID  string
	SellerID string
}

type OrderListInput struct {
	SellerID string
	Limit    int
	Offset   int
}

type OrderListItem struct {
	OrderID     string
	ProductName string
	Quantity    int
}

type CancelOrderInput struct {
	OrderID  string
	SellerID string
}
