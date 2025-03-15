package product

import (
	"github.com/vucongthanh92/go-test-exam/database"
	"gorm.io/gorm"

	"github.com/vucongthanh92/go-test-exam/internal/domain/interfaces"
)

type productCommandRepository struct {
	writeDb *gorm.DB
}

func NewProductCommandRepository(writeDb *database.GormWriteDb) interfaces.ProductCommandRepoI {
	return &productCommandRepository{
		writeDb: *writeDb,
	}
}
