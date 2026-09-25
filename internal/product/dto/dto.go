package dto

type FindManyQuery struct {
	Limit  int
	Offset int
}

type AddNewProductInput struct {
	SellerID string
	Name     string
	Stock    int
}

type BrowseProductsInput struct {
	Limit  int
	Offset int
}

type ProductItem struct {
	ID       string
	Name     string
}