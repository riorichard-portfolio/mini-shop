package dto

type FindAllBySellerIdQuery struct {
	SellerID string
	Limit    int
	Offset   int
}