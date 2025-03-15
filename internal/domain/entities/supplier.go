package entities

import "github.com/google/uuid"

type Supplier struct {
	ID   uuid.UUID `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Name string    `gorm:"column:reference;type:varchar(50);not null" json:"name"`
}

func (Supplier) TableName() string {
	return "suppliers"
}
