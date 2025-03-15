package product

import (
	"context"

	"github.com/vucongthanh92/go-base-utils/tracing"
	"github.com/vucongthanh92/go-test-exam/database"
	"gorm.io/gorm"

	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/internal/domain/entities"
	"github.com/vucongthanh92/go-test-exam/internal/domain/interfaces"
	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
)

type productQueryRepository struct {
	readDb *gorm.DB
}

func NewProductQueryRepository(readDb *database.GormReadDb) interfaces.ProductQueryRepoI {
	return &productQueryRepository{
		readDb: *readDb,
	}
}

func (repo *productQueryRepository) GetProductByFilter(ctx context.Context, filter models.ProductListFilter) (
	response []models.ProductListResponse, totalRows int64, errRes httpcommon.ErrorDTO) {

	ctx, span := tracing.StartSpanFromContext(ctx, "GetProductByFilter")
	defer span.End()

	tx := repo.readDb.WithContext(ctx).Model(&entities.Product{}).
		Select(filter.Columns).Count(&totalRows).
		Joins("LEFT JOIN categories ON products.category_id = categories.id").
		Joins("LEFT JOIN suppliers ON products.supplier_id = suppliers.id")

	if len(filter.Status) > 0 {
		tx = tx.Where("products.status IN ?", filter.Status)
	}

	tx = tx.Limit(filter.Limit).Offset(filter.Offset).Order("products.added_date DESC")

	err := tx.Find(&response).Error
	if err != nil {
		errRes.IsSystemError = true
		errRes.Error = err
		return response, totalRows, errRes
	}

	return response, totalRows, errRes
}
