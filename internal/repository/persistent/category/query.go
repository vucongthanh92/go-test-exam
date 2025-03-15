package category

import (
	"github.com/vucongthanh92/go-test-exam/database"
	"gorm.io/gorm"

	"github.com/vucongthanh92/go-test-exam/internal/domain/interfaces"
)

type categoryQueryRepository struct {
	readDb *gorm.DB
}

func NewCategoryQueryRepository(readDb *database.GormReadDb) interfaces.CategoryQueryRepoI {
	return &categoryQueryRepository{
		readDb: *readDb,
	}
}
