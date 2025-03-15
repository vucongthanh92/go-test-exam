package supplier

import (
	"github.com/vucongthanh92/go-test-exam/database"
	"gorm.io/gorm"

	"github.com/vucongthanh92/go-test-exam/internal/domain/interfaces"
)

type supplierQueryRepository struct {
	readDb *gorm.DB
}

func NewSupplierQueryRepository(readDb *database.GormReadDb) interfaces.SupplierQueryRepoI {
	return &supplierQueryRepository{
		readDb: *readDb,
	}
}
