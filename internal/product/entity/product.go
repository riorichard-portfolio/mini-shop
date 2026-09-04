package entity

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
) (Product, error) {
	if stock < 0 {
		return Product{}, InvalidStockErr
	}
	return Product{
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
		return InvalidQuantityErr
	}
	if p.stock < quantity {
		return InsufficientStockErr
	}
	p.stock -= quantity
	return nil
}
