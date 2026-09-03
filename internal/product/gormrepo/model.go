package gormrepo

type ProductGorm struct {
	ID       string `gorm:"column:id;type:uuid;primaryKey"`
	SellerID string `gorm:"column:seller_id;type:uuid"`
	Name     string `gorm:"column:name"`
	Stock    int    `gorm:"column:stock"`
}

func (ProductGorm) TableName() string {
	return "products"
}
