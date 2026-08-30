package gormrepo

type CustomerGorm struct {
	ID             string `gorm:"column:id;type:uuid;primaryKey"`
	Email          string `gorm:"column:email;type:varchar(255);not null;uniqueIndex"`
	HashedPassword string `gorm:"column:hashed_password;type:varchar(255);not null"`
}

func (CustomerGorm) TableName() string {
	return "customers"
}
