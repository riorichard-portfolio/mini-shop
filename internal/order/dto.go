package order

type MakeOrderInput struct {
	CustomerID  string
	ProductID   string
	Quantity    int
}

type CompleteOrderInput struct {
	OrderID string
	SellerID string
}

type FindAllBySellerIdQuery struct {
	SellerID string
	Limit int
	Offset int
}