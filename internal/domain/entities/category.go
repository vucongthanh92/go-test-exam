package entities

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Name string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
}

func (Category) TableName() string {
	return "categories"
}
