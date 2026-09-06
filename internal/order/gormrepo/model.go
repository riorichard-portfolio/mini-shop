package gormrepo

import "time"

type OrderGorm struct {
	ID          string    `gorm:"type:uuid;primaryKey"`
	CustomerID  string    `gorm:"type:uuid;not null"`
	ProductID   string    `gorm:"type:uuid;not null"`
	SellerID    string    `gorm:"type:uuid;not null"`
	ProductName string    `gorm:"type:varchar(255);not null"`
	Quantity    int       `gorm:"type:int;not null"`
	Status      string    `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null"`
}

func (OrderGorm) TableName() string {
	return "orders"
}