package entity

type Product struct {
	id          string
	sellerID    string
	name        string
	stockToDecr int
}

func NewProduct(
	id string,
	sellerID string,
	name string,
) (Product, error) {
	return Product{
		id:          id,
		sellerID:    sellerID,
		name:        name,
		stockToDecr: 0,
	}, nil
}

func (p *Product) ID() string       { return p.id }
func (p *Product) SellerID() string { return p.sellerID }
func (p *Product) Name() string     { return p.name }
func (p *Product) StockToDecr() int { return p.stockToDecr }

func (p *Product) DecreaseStock(quantity int, sellerID string) error {
	if p.sellerID != sellerID {
		return UnauthorizedSellerErr
	}
	if quantity <= 0 {
		return InvalidQuantityErr
	}
	p.stockToDecr += quantity
	return nil
}
