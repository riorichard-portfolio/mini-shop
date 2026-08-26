package entity

import "mini-shop/internal/pkg/bizerr"

type Product struct {
	id       string
	sellerID string
	name     string
	stock    int
}

func NewProduct(
	id string,
	sellerID string,
	name string,
	stock int,
) (*Product, error) {
	if stock < 0 {
		return nil, bizerr.New("INVALID_STOCK")
	}
	return &Product{
		id:       id,
		sellerID: sellerID,
		name:     name,
		stock:    stock,
	}, nil
}

func (p *Product) ID() string       { return p.id }
func (p *Product) SellerID() string { return p.sellerID }
func (p *Product) Name() string     { return p.name }
func (p *Product) Stock() int       { return p.stock }

func (p *Product) DecreaseStock(quantity int) error {
	if quantity <= 0 {
		return bizerr.New("INVALID_QUANTITY")
	}
	if p.stock < quantity {
		return bizerr.New("INSUFFICIENT_STOCK")
	}
	p.stock -= quantity
	return nil
}
