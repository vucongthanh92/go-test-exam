package entities

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID `gorm:"column:id;type:uuid;primaryKey"`
	Reference  string    `gorm:"column:reference;type:varchar(50);not null"`
	Name       string    `gorm:"column:name;type:varchar(255);not null"`
	AddedDate  time.Time `gorm:"column:added_date;type:date"`
	Status     string    `gorm:"column:status;type:varchar(50)"`
	CategoryID uuid.UUID `gorm:"column:category_id;type:uuid"`
	Price      float64   `gorm:"column:price;type:decimal(10,2);default:0"`
	StockCity  string    `gorm:"column:stock_city;type:varchar(100)"`
	SupplierID uuid.UUID `gorm:"column:supplier_id;type:uuid"`
	Quantity   int       `gorm:"column:quantity;type:int(4);default:0"`
}

func (Product) TableName() string {
	return "products"
}
