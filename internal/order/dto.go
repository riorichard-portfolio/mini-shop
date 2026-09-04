package order

type MakeOrderInput struct {
	CustomerID  string
	ProductID   string
	Quantity    int
}