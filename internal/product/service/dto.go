package service

type FindByIDOutput struct {
	ID       string
	Name     string
	SellerID string
}

type DecreaseStockInput struct {
	ID       string
	Quantity int
	SellerID string
}
