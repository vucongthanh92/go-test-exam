package category

import (
	"github.com/vucongthanh92/go-test-exam/database"
	"gorm.io/gorm"

	"github.com/vucongthanh92/go-test-exam/internal/domain/interfaces"
)

type categoryCommandRepository struct {
	writeDB *gorm.DB
}

func NewCategoryCommandRepository(writeDB *database.GormWriteDb) interfaces.CategoryCommandRepoI {
	return &categoryCommandRepository{
		writeDB: *writeDB,
	}
}
